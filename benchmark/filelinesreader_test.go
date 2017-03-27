package benchmark

import (
  "os"
  "testing"
  "bufio"
)

func BenchmarkFileLinesReader(b *testing.B) {
    f := os.Getenv("HOME") + "/crystal/src/class.cr"
    for i := 0; i < 500000; i++ {
        file, _ := os.Open(f)
        fileScanner := bufio.NewScanner(file)
        lineCount := 0
        for fileScanner.Scan() {
          lineCount++
        }
    }
}
