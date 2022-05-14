package os

import (
	"io/fs"
	"strings"
)

type OS interface {

	// 出力される実行ファイルの中に記載される。所有者を変更するコマンドの先頭に付与する文字列となる
	GetChownCd(owner string, group string, filePath string) string

	// 出力される実行ファイルの中に記載される。ディレクトリ作成をするコマンドの先頭に付与する文字列となる
	GetMkdirCd(filePath string) string

	// 出力される実行ファイルの中に記載される。実行権限を変更するコマンドの先頭に付与する文字列となる
	GetChmodCd(filePath string, permission fs.FileMode) string

	// 出力される実行ファイルの中に記載される。ファイルを作成するコマンドの先頭に付与する文字列となる
	GetMkfileCd(filePath string) string

	// 最終的に出力されるファイルの先頭行に出力される文字列。実行ファイルごとの予約語となる（例：sh → "#!/bin/bash\n\n"）
	GetExecuteFileReservedWord() string

	// ファイルのユーザとグループを抽出するコマンド
	GetExtractOwnerAndGroupCd() string

	// ファイルのユーザとグループを抽出するコマンドの引数
	GetExtractOwnerAndGroupCdArgs(string, fs.FileInfo) []string

	// 改行コードを取得する
	GetLineFeed() string

	// 指定された文字列の中から所有者情報を抽出する
	ExtractOwner(string) string

	// 指定された文字列の中から所有グループ情報を抽出する
	ExtractGroup(string) string
}

// この関数は、引数で指定された文字列の中に、2つ以上連続する空白があった場合に、空白が1つになるまで変換を行う。
func trimMultipleSpace(s string) string {
	for {
		// 2つ以上連続するスペースがなくなった場合は、ループを抜ける
		if !strings.Contains(s, "  ") {
			break
		}
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}
