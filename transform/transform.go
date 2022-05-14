package transform

import (
	"fmt"
	"path/filepath"

	"github.com/duplicate-dirs-go/extract"
	"github.com/duplicate-dirs-go/option"
	ddgos "github.com/duplicate-dirs-go/os"
)

type transformer struct {
	opt *option.Option
}

func NewTramsformer(opt *option.Option) *transformer {
	return &transformer{opt}
}

func (l *transformer) Transform(edArr []extract.ExtractData, o ddgos.OS) []string {

	result := []string{}

	// まずはおまじないを出力
	result = append(result, o.GetExecuteFileReservedWord())

	for _, ed := range edArr {

		// 対象ファイルの区切り文字として出力
		result = append(result, fmt.Sprintf("#target : %s%s", filepath.Join(ed.FileDirPath, ed.FileInfo.Name()), o.GetLineFeed()))

		// ディレクトリだった場合は、mkdirコマンドを書き出す
		if ed.FileInfo.IsDir() {
			result = append(result, fmt.Sprintf("%s%s", o.GetMkdirCd(filepath.Join(ed.FileDirPath, ed.FileInfo.Name())), o.GetLineFeed()))
		}

		// 対象がファイルだった場合は、touchコマンドを書き出す
		if !ed.FileInfo.IsDir() && !l.opt.OnlyDirLoad {
			result = append(result, fmt.Sprintf("%s%s", o.GetMkfileCd(filepath.Join(ed.FileDirPath, ed.FileInfo.Name())), o.GetLineFeed()))
		}

		// 所有者・グループの変更コマンドの生成
		if l.opt.NeedsOutputChown {
			result = append(result, fmt.Sprintf("%s%s", o.GetChownCd(ed.Owner, ed.Group, filepath.Join(ed.FileDirPath, ed.FileInfo.Name())), o.GetLineFeed()))
		}

		// 権限変更コマンドの生成
		if l.opt.NeedsOutputChmod {
			result = append(result, fmt.Sprintf("%s%s", o.GetChmodCd(filepath.Join(ed.FileDirPath, ed.FileInfo.Name()), ed.FileInfo.Mode().Perm()), o.GetLineFeed()))
		}

		// 改行コードを出力
		result = append(result, o.GetLineFeed())
	}

	return result
}
