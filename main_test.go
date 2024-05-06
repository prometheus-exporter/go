package main

import (
	"flag"
	"math/rand/v2"
	net_http "net/http"
	"os"
	"strconv"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/common-library/go/http"
)

func TestRun(t *testing.T) {
	address := ":" + strconv.Itoa(10000+rand.IntN(1000))
	timeout := strconv.Itoa(rand.IntN(10)+1) + "s"
	urlPath := "/url-path-" + strconv.Itoa(rand.IntN(10)+1)

	os.Args = []string{"test", "-address=" + address, "-timeout=" + timeout, "-url-path=" + urlPath}

	job := func() {
		if err := run(); err != nil {
			t.Fatal(err)
		}
	}

	test(t, address, urlPath, job)
}

func TestMain(t *testing.T) {
	address := ":" + strconv.Itoa(10000+rand.IntN(1000))
	timeout := strconv.Itoa(rand.IntN(10)+1) + "s"
	urlPath := "/url-path-" + strconv.Itoa(rand.IntN(10)+1)

	os.Args = []string{"test", "-address=" + address, "-timeout=" + timeout, "-url-path=" + urlPath}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	job := func() { main() }

	test(t, address, urlPath, job)
}

func test(t *testing.T, address, urlPath string, job func()) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()

		job()
	}()
	time.Sleep(100 * time.Millisecond)

	if response, err := http.Request("http://"+address+urlPath, net_http.MethodGet, nil, "", 5, "", "", nil); err != nil {
		t.Fatal(err)
	} else if response.StatusCode != net_http.StatusOK {
		t.Fatal("invalid -", response.StatusCode)
	}

	if err := syscall.Kill(os.Getpid(), syscall.SIGTERM); err != nil {
		t.Error(err)
	}

	wg.Wait()
}
