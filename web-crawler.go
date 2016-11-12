// https://go-tour-jp.appspot.com/concurrency/10

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var fetched map[string]bool

// -------- 動作の順序
// URLをフェッチする
// フェッチしたURLのリンクを探す
// リンクのURLをフェッチする

// -------- 問題点
// chを使いまわすと遅いフェッチがあったときに順番が崩れる
//   -> それによって再帰ごとに同期しないと使えない
// タイムアウト処理の実装?
//   -> ここでは考えないことにする

// チャネルの設計
//           root
// 1-1 1-2 1-3,        2-1 2-2 2-3
// 1-1-1 1-1-2, 1-2-1 1-2-2, ...
//   ---- 1-1 1-1-1 1-1-2 のようなグループでchを作っていく
//        1-1の1つ下の階層(1-1-x)が並列に処理できるようになる

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	var wg sync.WaitGroup
	ch := make(chan []string, 1)
	wg.Add(1)
	go _crawl(url, fetcher, ch, &wg)
	urls := <-ch

	// 残りのURLを指定の深さまでフェッチする
	wg.Add(1)
	go _fetch(urls, depth-1, fetcher, &wg)
	wg.Wait()

	return
}

func _fetch(urls []string, depth int, fetcher Fetcher, globalWg *sync.WaitGroup) {
	defer globalWg.Done()

	if depth <= 0 {
		return
	}

	var wg sync.WaitGroup
	fetchCh := make(chan []string, len(urls))
	for _, u := range urls {
		wg.Add(1)
		go _crawl(u, fetcher, fetchCh, &wg)
	}
	wg.Wait()
	close(fetchCh)

	// _crawlの結果を全て受信する
	for res := range fetchCh {
		globalWg.Add(1)
		go _fetch(res, depth-1, fetcher, globalWg)
	}
}

func _crawl(url string, fetcher Fetcher, ch chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("found: %s %q\n", url, body)
		ch <- urls
	}
}

func main() {
	fetched = make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// *それっぽく待つ
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
