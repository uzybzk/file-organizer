package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: file-organizer <directory>")
        return
    }
    
    directory := os.Args[1]
    
    err := organizeFiles(directory)
    if err != nil {
        fmt.Printf("Error organizing files: %v\n", err)
        return
    }
    
    fmt.Printf("Files organized successfully in %s\n", directory)
}

func organizeFiles(dir string) error {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return err
    }
    
    for _, file := range files {
        if !file.IsDir() {
            err := moveFileByExtension(dir, file.Name())
            if err != nil {
                return err
            }
        }
    }
    
    return nil
}

func moveFileByExtension(baseDir, filename string) error {
    ext := strings.ToLower(filepath.Ext(filename))
    if ext == "" {
        return nil
    }
    
    ext = strings.TrimPrefix(ext, ".")
    targetDir := filepath.Join(baseDir, ext)
    
    err := os.MkdirAll(targetDir, 0755)
    if err != nil {
        return err
    }
    
    oldPath := filepath.Join(baseDir, filename)
    newPath := filepath.Join(targetDir, filename)
    
    return os.Rename(oldPath, newPath)
}