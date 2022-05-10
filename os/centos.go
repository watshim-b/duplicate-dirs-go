package os

import (
	"fmt"
	"io/fs"
	"strings"
)

type centOs struct{}

func (o *centOs) GetChownCd(owner string, group string, filePath string) string {
	return fmt.Sprintf("%s %s:%s %s", "chown", owner, group, filePath)
}

func (o *centOs) GetMkdirCd(filePath string) string {
	return fmt.Sprintf("%s %s", "mkdir -p", filePath)
}

func (o *centOs) GetChmodCd(filePath string, permission fs.FileMode) string {
	return fmt.Sprintf("%s %o %s", "chmod", permission, filePath)
}

func (o *centOs) GetMkfileCd(filePath string) string {
	return fmt.Sprintf("%s %s", "touch", filePath)
}

func (o *centOs) GetExecuteFileReservedWord() string {
	return "#!/bin/bash"
}

func (o *centOs) GetLineFeed() string {
	return "\n"
}

func (o *centOs) GetExtractOwnerAndGroupCd() string {
	return "ls"
}

func (o *centOs) GetExtractOwnerAndGroupCdArgs(filePath string, fileInfo fs.FileInfo) []string {
	if fileInfo.IsDir() {
		return []string{"-lda", filePath}
	} else {
		return []string{"-la", filePath}
	}
}

func (o *centOs) ExtractOwner(strArr string) string {
	return strings.Split(trimMultipleSpace(strArr), " ")[2]
}

func (o *centOs) ExtractGroup(strArr string) string {
	return strings.Split(trimMultipleSpace(strArr), " ")[3]
}
