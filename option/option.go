package option

import (
	"flag"
)

type Option struct {

	// 最終的に出力する時のファイル名を指定する
	LoadFileNm string

	// shにchmodコマンドを出力する必要があるかどうかの判定
	NeedsOutputChmod bool

	// ディレクトリのみ出力したいかどうかを指定する
	OnlyDirLoad bool

	// shにchownコマンドを出力する必要があるかどうかの判定
	NeedsOutputChown bool

	// 復元対象としたい場所のトップディレクトリの指定
	TopDir string

	// 利用するOSを指定する。指定されたOSによって処理が後続の処理が柔軟に対応できる
	Os string

	// スキップしたいディレクトリ名を指定する。（あいまい検索で指定された文字列を含むディレクトリは、復元対象からスキップされる）
	// SkipDir []string
}

func (o *Option) BindFromFlag() error {
	flag.StringVar(&o.LoadFileNm, "lfn", "duplicate", "generate file name.")
	flag.StringVar(&o.TopDir, "td", "/home", "file licursive top dir.")
	flag.BoolVar(&o.NeedsOutputChown, "aco", false, "need ouput chown command.")
	flag.BoolVar(&o.NeedsOutputChmod, "acm", false, "need ouput chmod command.")

	//　実行環境を定義するための変数
	osStr := flag.String("os", "", "specific target os.")

	// 初期値がtrueの場合にBoolVarが使えないので、こちらに変更
	odl := flag.Bool("odl", true, "output only directory related command.")

	flag.Parse()
	o.OnlyDirLoad = *odl
	o.Os = *osStr

	println("")
	println("//////////////////////////////")
	println("/////////args details/////////")
	println("//////////////////////////////")
	println("・ofn:" + o.LoadFileNm)
	println("・td:" + o.TopDir)
	print("・odl:")
	println(o.OnlyDirLoad)
	print("・aco:")
	println(o.NeedsOutputChown)
	print("・acm:")
	println(o.NeedsOutputChmod)
	print("・os:")
	println(*osStr)
	println("")

	return nil
}
