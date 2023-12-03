package fundamentals

type Reader interface {
	Read([]byte) (int, error)
}

type File struct{}

func (f File) Read(b []byte) (n int, err error)

var f File
var r Reader = f
var f2 File = r.(File)
var f3, ok = r.(File)
