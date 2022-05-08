package loader

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/duplicate-dirs-go/extract"
	"github.com/duplicate-dirs-go/option"
	ddgos "github.com/duplicate-dirs-go/os"
)

type loader struct {
	opt *option.Option
}

func NewLoader(opt *option.Option) *loader {
	return &loader{opt}
}

func (l *loader) Load(w io.Writer, edArr []extract.ExtractData, c ddgos.Commader) error {
	// まずはおまじないを出力
	_, err := w.Write([]byte(c.GetExecuteFileReservedWord()))
	if err != nil {
		return err
	}

	for _, ed := range edArr {

		// 対象ファイルの区切り文字として出力
		_, err = w.Write([]byte(fmt.Sprintf("#target : %s%s", filepath.Join(ed.FileDirPath, ed.FileInfo.Name()), c.GetLineFeed())))
		if err != nil {
			return err
		}

		// ディレクトリだった場合は、mkdirコマンドを書き出す
		if ed.FileInfo.IsDir() {
			_, err = w.Write([]byte(fmt.Sprintf("%s%s", c.GetMkdirCd(filepath.Join(ed.FileDirPath, ed.FileInfo.Name())), c.GetLineFeed())))
			if err != nil {
				return err
			}
		}

		// 対象がファイルだった場合は、touchコマンドを書き出す
		if !ed.FileInfo.IsDir() && !l.opt.OnlyDirLoad {
			_, err := w.Write([]byte(fmt.Sprintf("%s%s", c.GetMkfileCd(filepath.Join(ed.FileDirPath, ed.FileInfo.Name())), c.GetLineFeed())))
			if err != nil {
				return err
			}
		}

		// 所有者・グループの変更コマンドの生成
		if l.opt.NeedsOutputChown {
			_, err := w.Write([]byte(fmt.Sprintf("%s%s", c.GetChownCd(ed.Owner, ed.Group, filepath.Join(ed.FileDirPath, ed.FileInfo.Name())), c.GetLineFeed())))
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
		}

		// 権限変更コマンドの生成
		if l.opt.NeedsOutputChmod {
			_, err := w.Write([]byte(fmt.Sprintf("%s%s", c.GetChmodCd(filepath.Join(ed.FileDirPath, ed.FileInfo.Name()), ed.FileInfo.Mode().Perm()), c.GetLineFeed())))
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
		}

		// 改行コードを出力
		_, err = w.Write([]byte(c.GetLineFeed()))
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

	}
	return nil
}
