## Reference

- [func DecodeTo](#funcDecodeTo)
- [func Download](#funcDownload)
- [func Unzip](#funcUnzip)

## Usage

\
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
<a name='funcDownload'></a>**func Download**

```go
 Download(URL, destFolder, destFileName string) error
```

<details>

<summary>Downloads from a given URL to a destination file. The destination file can be specified or left blank (in which case it will be automatically detected).</summary>

Downloads a file from a URL and saves it to a local file. Download progress is reported in bytes to stdout. If possible, the content length will be used to calculate the percentage completion.

If an empty destFileName string is provided, the file name will be calculated from the base of the provided URL.

</details>

\
<a name='funcUnzip'></a>**func Unzip**

```go
func Unzip(source, destination string) ([]string, error)
```

<details>

<summary>Unzips a .zip file to the destination, retaining nested structure.</summary>

Unzip extracts a .zip archive into the specified destination directory. It returns a string slice of the filepaths for each extracted file, relative to the destination root.

If any of the destination files or folders already exists, they will be overwritten assuming the process has write permissions to the destination.

The destination parent will be named after the input .zip file.
e.g. if the input .zip file is named "test.zip", and the destination is "/data/project", the final destination will be "/data/project/test".

</details>
