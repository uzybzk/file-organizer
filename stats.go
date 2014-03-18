package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "strings"
)

type Stats struct {
    TotalFiles int
    Extensions map[string]int
}

func collectStats(dir string) (*Stats, error) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return nil, err
    }
    
    stats := &Stats{
        TotalFiles: 0,
        Extensions: make(map[string]int),
    }
    
    for _, file := range files {
        if !file.IsDir() {
            stats.TotalFiles++
            ext := strings.ToLower(filepath.Ext(file.Name()))
            if ext != "" {
                ext = strings.TrimPrefix(ext, ".")
                stats.Extensions[ext]++
            }
        }
    }
    
    return stats, nil
}

func printStats(stats *Stats) {
    fmt.Printf("Found %d files\n", stats.TotalFiles)
    fmt.Println("Extensions:")
    for ext, count := range stats.Extensions {
        fmt.Printf("  .%s: %d files\n", ext, count)
    }
}