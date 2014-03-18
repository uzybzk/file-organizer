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
        fmt.Println("       file-organizer --help")
        return
    }
    
    if os.Args[1] == "--help" || os.Args[1] == "-h" {
        showHelp()
        return
    }
    
    directory := os.Args[1]
    
    fmt.Printf("Organizing files in: %s\n", directory)
    
    // Show stats before organizing
    stats, err := collectStats(directory)
    if err != nil {
        fmt.Printf("Error collecting stats: %v\n", err)
        return
    }
    printStats(stats)
    fmt.Println()
    
    err = organizeFiles(directory)
    if err != nil {
        fmt.Printf("Error organizing files: %v\n", err)
        return
    }
    
    fmt.Printf("Files organized successfully in %s\n", directory)
}

func showHelp() {
    fmt.Println("File Organizer - Organize files by extension")
    fmt.Println()
    fmt.Println("Usage:")
    fmt.Println("  file-organizer <directory>")
    fmt.Println()
    fmt.Println("Examples:")
    fmt.Println("  file-organizer ~/Downloads")
    fmt.Println("  file-organizer /path/to/messy/folder")
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