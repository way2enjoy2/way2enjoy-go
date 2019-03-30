# Way2Enjoy API client for Golang

------

Golang client for the Way2Enjoy API, used for [way2Enjoy](https://way2Enjoy.com) . Way2Enjoy compresses or resize your images intelligently. Read more at [http://way2Enjoy.com](http://way2Enjoy.com).

## Documentation

[Go to the documentation for the HTTP client](https://way2Enjoy.com/developers).

## Installation

Install the API client with `go get`.

```shell
go get -u github.com/way2enjoy2/way2Enjoy-go
```

## Usage

- About key

    Get your API key from  https://way2Enjoy.com/developers

- Compress
    ```golang
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
    ```

- Resize
    ```golang
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
    ```

- ***Notice:***

    Way2Enjoy.ResizeMethod support `scale`, `fit` and `cover`. If used fit or cover, you must provide `both a width and a height`. But used scale, you must provide either a target width or a target height, `but not both`.


- More usage, please see [way2Enjoy_test.go](./way2Enjoy_test.go)

## Running tests

```shell
cd $GOPATH/src/github.com/way2Enjoy2/way2Enjoy-go
go test
```

## License

This software is licensed under the MIT License. [View the license](LICENSE).
