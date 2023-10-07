package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:children`
}

var stRootDir string
var stSeparator string
var iRootNode Node

const stJsonFileName = "dir.json"

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]
	gnJsonFileBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)

	fmt.Println(stWorkDir+stJsonFileName, gnJsonFileBytes)
	err := json.Unmarshal(gnJsonFileBytes, &iRootNode)

	if err != nil {
		panic("error")
	}
}

func parseNode(iNode Node, stParentDir string) {
	if iNode.Text != "" {
		createDir(iNode, stParentDir)
	}

	if stParentDir != "" {
		stParentDir = stParentDir + stSeparator
	}

	if iNode.Text != "" {
		stParentDir = stParentDir + stSeparator + iNode.Text
	}

	for _, v := range iNode.Children {
		parseNode(v, stParentDir)
	}
}

func createDir(iNode Node, stParentDir string) {
	stDirPath := stRootDir + stSeparator

	if stParentDir != "" {
		stDirPath = stDirPath + stParentDir + stSeparator
	}

	stDirPath = stDirPath + iNode.Text
	err := os.MkdirAll(stDirPath, fs.ModePerm)
	if err != nil {
		panic("error")
	}
}

func TestDir(t *testing.T) {
	loadJson()
	parseNode(iRootNode, "")
}
