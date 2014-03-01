# File Organizer

A CLI tool to organize files by extension into subdirectories.

## Usage

```bash
go run main.go /path/to/messy/directory
```

## What it does

Organizes files by their extensions:
- `.txt` files go to `txt/` folder
- `.jpg` files go to `jpg/` folder
- etc.

## Example

Before:
```
downloads/
  document.pdf
  photo.jpg
  song.mp3
```

After:
```
downloads/
  pdf/document.pdf
  jpg/photo.jpg
  mp3/song.mp3
```