# Benchmark of various http file send method 

## Rapid Conclusion
io.pipe is good, but not fast enough when dealing with large files.

If you need more speed with large files, consider using zero copy methods(os.splice, os.pipe, etc.)

Note: I excluded sendFile method of net package because it is hard to use with http package.

## Methods

1. Use os pipe

```go
func osPipeSend(fileName, url string) {
	r, w, _ := os.Pipe()
	writer := multipart.NewWriter(w)

	file, _ := os.Open(fileName)
	defer file.Close()
	go func() {
		writer.CreateFormFile("file", filepath.Base(file.Name()))
		w.ReadFrom(file)
		writer.Close()
		w.Close()
	}()

	http.Post(url, writer.FormDataContentType(), r)
}
```

2. Use zerocopy library

```go
type readCloserPipeWrap struct {
	*zerocopy.Pipe
}

func (p *readCloserPipeWrap) Close() error {
	return p.CloseRead()
}

func zeroCopyLibSend(fileName, url string) {
	p, _ := zerocopy.NewPipe()
	rw := &readCloserPipeWrap{p}
	writer := multipart.NewWriter(rw)

	file, _ := os.Open(fileName)
	defer file.Close()
	go func() {
		writer.CreateFormFile("file", filepath.Base(file.Name()))
		rw.ReadFrom(file)
		writer.Close()
		rw.CloseWrite()
	}()

	http.Post(url, writer.FormDataContentType(), rw)
}
```

3. Use io pipe

```go
func ioPipeSend(fileName, url string) {
	r, w := io.Pipe()
	writer := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		file, _ := os.Open(fileName)
		defer file.Close()
		part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
		io.Copy(part, file)
		writer.Close()
	}()

	http.Post(url, writer.FormDataContentType(), r)
}
```

4. No pipe. just enjoy memory party, memory is cheaper than programmer.

```go
func naiveSend(fileName, url string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	http.Post(url, writer.FormDataContentType(), body)
}
```

## Benchmark command
setup shell
``` sh
go run setup.go
```

benchmark shell (optionally limit benchtime to prevent port exhaustion)
```sh
go test -benchmem . -bench . -benchtime=10x
```

## Benchmark result

> **Warning**
> My computer environment is very noising, please focus on using memory (B/op and allocs/op) rather than naive process performance

Some prettier works(column name, indentation, etc.) on result

### Windows
```log
goos: windows
goarch: amd64
pkg: example
cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz

benchmark name                   count       ns/op      B/op  allocs/op
BenchmarkOsPipeSend100MB-8         567     2060950     87499        181
BenchmarkZeroCopyLibSend100MB-8    532     2204932     87588        187
BenchmarkIoPipeSend100MB-8         441     2583869     86977        293
BenchmarkNaiveSend100MB-8           12   117695575 335588924        179

BenchmarkOsPipeSend10MB-8          594     2077345     87567        182
BenchmarkZeroCopyLibSend10MB-8     590     2269528     87591        186
BenchmarkIoPipeSend10MB-8          450     3265677     86872        294
BenchmarkNaiveSend10MB-8            73    14378955  41987294        175

BenchmarkOsPipeSend1MB-8           580     1891629     87548        180
BenchmarkZeroCopyLibSend1MB-8      596     1899646     87618        185
BenchmarkIoPipeSend1MB-8           679     1671816     86756        200
BenchmarkNaiveSend1MB-8            512     2335895   2666803        172

BenchmarkOsPipeSend100KB-8         686     1726934     87390        171
BenchmarkZeroCopyLibSend100KB-8    633     1849858     87404        175
BenchmarkIoPipeSend100KB-8         724     1634127     86382        171
BenchmarkNaiveSend100KB-8          753     1590394    373179        168

BenchmarkOsPipeSend10KB-8          814     1549107     87387        168
BenchmarkZeroCopyLibSend10KB-8     867     1267104     87471        172
BenchmarkIoPipeSend10KB-8          987     1465305     86354        168
BenchmarkNaiveSend10KB-8           862     1365444     71313        166

BenchmarkOsPipeSend1KB-8           928     1262907     87386        168
BenchmarkZeroCopyLibSend1KB-8      931     1279530     87455        172
BenchmarkIoPipeSend1KB-8           940     1253947     86365        168
BenchmarkNaiveSend1KB-8            907     1290179     54718        164
```

### Linux (WSL: Ubuntu-20.04)
```log
goos: linux
goarch: amd64
pkg: example
cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz

benchmark name                   count       ns/op      B/op  allocs/op
BenchmarkOsPipeSend100MB-8        1623      711531     83852        161
BenchmarkZeroCopyLibSend100MB-8   2119      666922     52056        177
BenchmarkIoPipeSend100MB-8         163     9662024     83720        244
BenchmarkNaiveSend100MB-8           13   166625393 335586727        162

BenchmarkOsPipeSend10MB-8         1250      864626     83850        163
BenchmarkZeroCopyLibSend10MB-8    1501     1071810     52218        177
BenchmarkIoPipeSend10MB-8          195    10791563     83770        252
BenchmarkNaiveSend10MB-8           110    24652351  41984802        155

BenchmarkOsPipeSend1MB-8          1477      826677     83876        165
BenchmarkZeroCopyLibSend1MB-8     1671      919370     52114        177
BenchmarkIoPipeSend1MB-8          1310     2139825     84334        178
BenchmarkNaiveSend1MB-8            835     1213371   2663698        151

BenchmarkOsPipeSend100KB-8        1602      767052     83690        148
BenchmarkZeroCopyLibSend100KB-8   1508      705544     52139        177
BenchmarkIoPipeSend100KB-8        2054      540105     83964        149
BenchmarkNaiveSend100KB-8         2402      507276    370370        148

BenchmarkOsPipeSend10KB-8         2557      450030     83680        145
BenchmarkZeroCopyLibSend10KB-8    3129      444820     51382        167
BenchmarkIoPipeSend10KB-8         2810      428353     83995        146
BenchmarkNaiveSend10KB-8          3105     1910439     68452        146

BenchmarkOsPipeSend1KB-8          3102      341200     83696        145
BenchmarkZeroCopyLibSend1KB-8     3549      365182     51442        167
BenchmarkIoPipeSend1KB-8          3411      542196     83985        146
BenchmarkNaiveSend1KB-8            450     2593385     52038        142
```
