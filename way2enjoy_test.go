package main

import (
	"testing"

	"io/ioutil"

	"github.com/way2Enjoy2/way2Enjoy-go/way2Enjoy"
)

const Key = "rcPZm3Zrg_1DbjYtV6AXM_-53Jg9wuWB"

func TestCompressFromFile(t *testing.T) {
	Way2Enjoy.SetKey(Key)
	source, err := Way2Enjoy.FromFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	err = source.ToFile("./test_output/CompressFromFile.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Compress successful")
}

func TestCompressFromBuffer(t *testing.T) {
	Way2Enjoy.SetKey(Key)

	buf, err := ioutil.ReadFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	source, err := Way2Enjoy.FromBuffer(buf)
	if err != nil {
		t.Error(err)
		return
	}

	err = source.ToFile("./test_output/CompressFromBuffer.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Compress successful")
}

func TestCompressFromUrl(t *testing.T) {
	Way2Enjoy.SetKey(Key)
	url := "http://pic.tugou.com/realcase/1481255483_7311782.jpg"
	source, err := Way2Enjoy.FromUrl(url)
	if err != nil {
		t.Error(err)
		return
	}
	err = source.ToFile("./test_output/CompressFromUrl.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Compress successful")
}

func TestResizeFromFile(t *testing.T) {
	Way2Enjoy.SetKey(Key)
	source, err := Way2Enjoy.FromFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	err = source.Resize(&Way2Enjoy.ResizeOption{
		Method: Way2Enjoy.ResizeMethodFit,
		Width:  100,
		Height: 100,
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = source.ToFile("./test_output/ResizeFromFile.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Resize successful")
}

func TestResizeFromBuffer(t *testing.T) {
	Way2Enjoy.SetKey(Key)

	buf, err := ioutil.ReadFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	source, err := Way2Enjoy.FromBuffer(buf)
	if err != nil {
		t.Error(err)
		return
	}

	err = source.Resize(&Way2Enjoy.ResizeOption{
		Method: Way2Enjoy.ResizeMethodScale,
		Width:  200,
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = source.ToFile("./test_output/ResizesFromBuffer.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Resize successful")
}

func TestResizeFromUrl(t *testing.T) {
	Way2Enjoy.SetKey(Key)
	url := "http://pic.tugou.com/realcase/1481255483_7311782.jpg"
	source, err := Way2Enjoy.FromUrl(url)
	if err != nil {
		t.Error(err)
		return
	}

	err = source.Resize(&Way2Enjoy.ResizeOption{
		Method: Way2Enjoy.ResizeMethodCover,
		Width:  300,
		Height: 100,
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = source.ToFile("./test_output/ResizeFromUrl.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Resize successful")
}
