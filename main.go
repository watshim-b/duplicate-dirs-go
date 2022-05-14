package main

import (
	"fmt"
	"os"

	"github.com/duplicate-dirs-go/extract"
	"github.com/duplicate-dirs-go/loader"
	"github.com/duplicate-dirs-go/option"
	ddgos "github.com/duplicate-dirs-go/os"
	"github.com/duplicate-dirs-go/transform"
)

func main() {

	println("【info】 start process.")

	// 指定された引数を変数に割り当てる
	opt := &option.Option{}
	opt.BindFromFlag()

	// 対象のosを特定して構造体を初期化する
	osKind := ddgos.ValueOf(opt.Os)
	if osKind == ddgos.None {
		println(fmt.Sprintf("指定されたOSは存在しません。 利用可能なOSは、 %s です。", ddgos.AvailableOS()))
		os.Exit(1)
	}
	opSys := osKind.GenerateOSInstance()

	// 復元元のファイル情報を抽出する
	e := extract.NewExtractor(opt)
	extractDataArr, err := e.ExtractFilePath(opSys)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	// 権限情報も抽出する櫃よぐある場合は、一緒に抽出する
	if opt.NeedsOutputChown {
		extractDataArr, err = e.ExtractOwnerAndGroup(extractDataArr, opSys)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}

	// 抽出したデータを出力用にに変換する
	t := transform.NewTramsformer(opt)
	transformDataArr := t.Transform(extractDataArr, opSys)

	// 出力用のファイルを生成する
	f, err := os.OpenFile(opt.LoadFileNm, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	// 最終的なファイルを出力する
	l := loader.NewLoader(opt)
	err = l.Load(f, transformDataArr)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println("【info】 end process.")
}
