package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/saasan/go-s2dir"
	"github.com/saasan/go-s2file"
	"github.com/saasan/go-term"
)

func renameFileToFolderName(parent string, dirname string) error {
	// サブディレクトリとファイルを列挙
	dirs, files, err := s2dir.Read(dirname)
	if err != nil {
		return err
	}

	if len(dirs) != 0 || len(files) != 1 {
		return nil
	}

	// 最終的な移動先
	basename := filepath.Base(dirname)
	ext := filepath.Ext(files[0].Name())
	newpath := filepath.Join(parent, basename+ext)

	if s2file.Exists(newpath) {
		return nil
	}

	// ファイルパスの長さ対策のため一旦親フォルダへ移動
	oldpath := filepath.Join(dirname, files[0].Name())
	parentpath := filepath.Join(parent, files[0].Name())
	s2file.Rename(oldpath, parentpath)

	if isEmpty, err := s2dir.IsEmpty(dirname); err != nil {
		return err
	} else if !isEmpty {
		return nil
	}

	// ディレクトリが空であれば削除
	fmt.Printf("空フォルダを削除: %s\n", dirname)
	if err := os.Remove(dirname); err != nil {
		return fmt.Errorf("%s を削除できません。: %w", dirname, err)
	}

	// ファイル名をフォルダの名前へ変更
	s2file.Rename(parentpath, newpath)

	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("処理対象: %s\n", arg)

		// サブディレクトリを列挙
		dirs, _, err := s2dir.Read(arg)
		if err != nil {
			continue
		}

		// サブディレクトリ内のファイルを移動
		for _, dir := range dirs {
			path := filepath.Join(arg, dir.Name())
			if err := renameFileToFolderName(arg, path); err != nil {
				fmt.Println(err)
			}
		}
	}

	if runtime.GOOS == "windows" {
		fmt.Println("Press any key to continue...")
		term.WaitAnyKey()
	}
}
