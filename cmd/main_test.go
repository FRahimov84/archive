package main

import (
	"sync"
	"testing"
)

/// competitive 3,511720497289684 times faster than consistently Archiver) for 100 files
/// competitive 3,666694080649804 times faster than consistently Archiver) for 200 files

func Benchmark_competitiveArchiver(b *testing.B) {
	wg := sync.WaitGroup{}
	strings := []string{"3601892.png", "course.pdf", "SteamSetup.exe", "video.mp4"}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wgg *sync.WaitGroup) {
			defer wgg.Done()
			competitiveArchiver(strings)
		}(&wg)
	}
	wg.Wait()
}

func Benchmark_consistentlyArchiver(b *testing.B) {
	strings := []string{"3601892.png", "course.pdf", "SteamSetup.exe", "video.mp4"}
	for i := 0; i < 10; i++ {
		consistentlyArchiver(strings)
	}
}