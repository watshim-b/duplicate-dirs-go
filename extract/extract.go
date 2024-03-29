package extract

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/duplicate-dirs-go/option"
	ddgos "github.com/duplicate-dirs-go/os"
)

type extractor struct {
	opt *option.Option
}

type ExtractData struct {
	FileInfo    os.FileInfo
	FileDirPath string
	Owner       string
	Group       string
}

func NewExtractor(opt *option.Option) *extractor {
	return &extractor{opt}
}

// TODO ExtractOwnerAndGroup・ExtractFilePathをトランスフォーマーに食わせる構成に変更したい。

// フォルダを走査して、ファイルパスとファイル情報を抽出する
func (e *extractor) ExtractFilePath(o ddgos.OS) ([]ExtractData, error) {

	edArr := []ExtractData{}

	// topdirの情報を取得する
	td, err := os.Open(e.opt.TopDir)
	defer td.Close()

	if err != nil {
		return nil, err
	}

	fi, err := td.Stat()
	if err != nil {
		return nil, err
	}

	edArr = append(edArr, ExtractData{FileInfo: fi, FileDirPath: filepath.Dir(e.opt.TopDir)})

	rtEdArr, err := e.extractRecursiveFile(e.opt.TopDir, o)
	if err != nil {
		return nil, err
	}
	edArr = append(edArr, rtEdArr...)

	return edArr, nil
}

// 再帰的にフォルダを走査して、ファイルパスとファイル情報を抽出する
func (e *extractor) extractRecursiveFile(topDir string, o ddgos.OS) ([]ExtractData, error) {

	edArr := []ExtractData{}

	// topdir以下のファイルの情報を取得する
	files, err := ioutil.ReadDir(topDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		if file.IsDir() {
			edArr = append(edArr, ExtractData{FileInfo: file, FileDirPath: topDir})
			rtEdArr, err := e.extractRecursiveFile(filepath.Join(topDir, file.Name()), o)
			if err != nil {
				return nil, err
			}
			edArr = append(edArr, rtEdArr...)
			continue
		}

		// ディレクトリだけを出力する場合は、ファイルの情報は格納しないので、スキップ
		if e.opt.OnlyDirLoad {
			continue
		}

		edArr = append(edArr, ExtractData{FileInfo: file, FileDirPath: topDir})
	}

	return edArr, nil
}

// 所有者とグループ情報を抽出する
func (e *extractor) ExtractOwnerAndGroup(edArr []ExtractData, o ddgos.OS) ([]ExtractData, error) {
	for loopN, ed := range edArr {

		// 以上をして最終的にアウトプットして、shファイルが作成される。

		var filePath string
		if ed.FileInfo.IsDir() {
			filePath = ed.FileDirPath
		} else {
			filePath = filepath.Join(ed.FileDirPath, ed.FileInfo.Name())
		}

		result, err := exec.Command(o.GetExtractOwnerAndGroupCd(), o.GetExtractOwnerAndGroupCdArgs(filePath, ed.FileInfo)...).Output()

		// エラーがあった場合は処理を終了する
		if err != nil {
			return nil, err
		}

		// ユーザーとグループの設定
		edArr[loopN].Owner = o.ExtractOwner(string(result))
		edArr[loopN].Group = o.ExtractGroup(string(result))
	}
	return edArr, nil
}
