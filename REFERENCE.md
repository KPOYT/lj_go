## Reference

- [func DecodeTo](#funcDecodeTo)
- [func Download](#funcDownload)
- [func Unzip](#funcUnzip)

## Usage

<a name='funcDecodeTo'></a>**func DecodeTo**

```go
DecodeTo(i interface{}, jsonPath string, maxReadSize int64)
```

Decodes a json file into a corresponding json struct, disallowing missing fields.

<a name='funcDownload'></a>**func Download**

```go
 Download(URL, destFolder, destFileName string) error
```

Downloads from a given URL to a destination file. The destination file can be specified or left blank (in which case it will be automatically detected).

<a name='funcUnzip'></a>**func Unzip**

```go
func Unzip(source, destination string) ([]string, error)
```

Unzips a .zip file to the destination, retaining nested structure.
