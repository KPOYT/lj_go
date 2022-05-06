## Reference

- [func DecodeTo](#funcDecodeTo)
- [func Download](#funcDownload)
- [func Unzip](#funcUnzip)

## Usage

<a name='funcDecodeTo'></a>**func DecodeTo**

```go
DecodeTo(i interface{}, jsonPath string, maxReadSize int64)
```

<details>

<summary>Decodes a json file into a corresponding json struct, disallowing missing fields.</summary>

Decodes a json file to an interface. Interface fields must match the json to be decoded. Extra fields which are not included in the json file will be ignored.

If the json file contains multiple objects, only the first one will be decoded.

If the json file is empty, an error will be returned.

maxReadSize is the maximum size in bytes of the json file to read.
If the file is larger than this, an error is returned. Set to 0 to read the whole file regardless of its size.

</details>

\
\
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
