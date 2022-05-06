## A collection of useful functions I've collected over time which I reuse often.

## How to use

```bash
go get github.com/lambdajack/lj_go
```

```go
import "github.com/lambdajack/lj_go/pkg/lj_archive"

func main() {
	files, err := lj_archive.Unzip("path/to/file.zip", "path/to/out/dir")
}
```

## Reference

- [func DecodeTo](REFERENCE.md#funcDecodeTo)
- [func Download](REFERENCE.md#funcDownload)
- [func Unzip](REFERENCE.md#funcUnzip)

### Requirements

In order to be included, each function must:

1. Be stable. Where breaking changes are made, use versioning or just a completeley new function.
2. Be testable and have such tests included here.
3. Try to be as environment agnostic _as possible_ (i.e. as useful in a web api as it is in a cli application). Sometimes this is not possible and where it is not, usually multiple versions of the same function are available for each environment.
4. Do only one thing.
