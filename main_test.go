package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/saasan/go-s2file"
)

// 正常系
func TestRenameFileToFolderNameNormal(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestRenameFileToFolderName")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1_1  = filepath.Join(tmpDir, "dir1_1")
		file1_2 = filepath.Join(tmpDir, "dir1_1", "file1_2.txt")
		move1_2 = filepath.Join(tmpDir, "dir1_1.txt")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir1_1, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file1_2, []byte("file1_2"), os.ModePerm)

	renameFileToFolderName(tmpDir, dir1_1)

	if s2file.Exists(dir1_1) {
		t.Error("dir1_1失敗")
	}
	if s2file.Exists(file1_2) {
		t.Error("file1_2失敗")
	}
	if !s2file.Exists(move1_2) {
		t.Error("move1_2失敗")
	}
}

// 親フォルダに同名のファイルが存在する場合
func TestRenameFileToFolderNameExist(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestRenameFileToFolderName")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1_1  = filepath.Join(tmpDir, "dir1_1")
		file1_2 = filepath.Join(tmpDir, "dir1_1", "file1_2.txt")
		move1_2 = filepath.Join(tmpDir, "dir1_1.txt")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir1_1, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file1_2, []byte("file1_2"), os.ModePerm)
	os.WriteFile(move1_2, []byte("move1_2"), os.ModePerm)

	renameFileToFolderName(tmpDir, dir1_1)

	if !s2file.Exists(dir1_1) {
		t.Error("dir1_1失敗")
	}
	if !s2file.Exists(file1_2) {
		t.Error("file1_2失敗")
	}
	if !s2file.Exists(move1_2) {
		t.Error("move1_2失敗")
	}
}

// ファイルの拡張子がない場合
func TestRenameFileToFolderNameDuplicate(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestRenameFileToFolderName")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1_1  = filepath.Join(tmpDir, "dir1_1")
		file1_2 = filepath.Join(tmpDir, "dir1_1", "file1_2")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir1_1, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file1_2, []byte("file1_2"), os.ModePerm)

	renameFileToFolderName(tmpDir, dir1_1)

	if !s2file.Exists(dir1_1) {
		t.Error("dir1_1失敗")
	}
	if !s2file.Exists(file1_2) {
		t.Error("file1_2失敗")
	}
}

// ファイルが複数ある場合
func TestRenameFileToFolderNameMultiple(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestRenameFileToFolderName")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1_1    = filepath.Join(tmpDir, "dir1_1")
		file1_2_1 = filepath.Join(tmpDir, "dir1_1", "file1_2_1.txt")
		file1_2_2 = filepath.Join(tmpDir, "dir1_1", "file1_2_2.txt")
		move1_2   = filepath.Join(tmpDir, "dir1_1.txt")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir1_1, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file1_2_1, []byte("file1_2"), os.ModePerm)
	os.WriteFile(file1_2_2, []byte("file1_2"), os.ModePerm)

	renameFileToFolderName(tmpDir, dir1_1)

	if !s2file.Exists(dir1_1) {
		t.Error("dir1_1失敗")
	}
	if !s2file.Exists(file1_2_1) {
		t.Error("file1_2_1失敗")
	}
	if !s2file.Exists(file1_2_2) {
		t.Error("file1_2_2失敗")
	}
	if s2file.Exists(move1_2) {
		t.Error("move1_2失敗")
	}
}
