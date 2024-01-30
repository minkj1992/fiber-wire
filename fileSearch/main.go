package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/k0kubun/pp/v3"
)

var (
	matches   []string
	wg        = sync.WaitGroup{}
	lock      = sync.Mutex{}
	maxLength = 2048
)

func setPP() {
	pp.SetColorScheme(pp.ColorScheme{
		Integer: pp.Green | pp.Bold,
		Float:   pp.Black | pp.BackgroundWhite | pp.Bold,
		String:  pp.Yellow,
	})
	pp.BufferFoldThreshold = maxLength
}

// /Users/minwook/code
func baseDir() (parent string) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(cwd)
}

func fileSearch(root, target string) {
	pp.Printf("Searching in %s\n", root)

	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), target) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), target)
		}

	}

	wg.Done()
}

func main() {
	dir := baseDir()
	setPP()
	pp.Printf("Base directory is %s\n", dir)

	wg.Add(1)
	go fileSearch(dir, "README.md")
	wg.Wait()

	pp.Printf("%v matched\n", len(matches))
	pp.Printf("%+v\n", matches)
}
