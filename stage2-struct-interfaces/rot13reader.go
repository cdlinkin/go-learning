/*
exercise:

Write your own io.Reader (for example, rot13Reader) and connect it via io.Copy(os.Stdout, r).
*/
package stage2

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 'a' + (p[i]-'a'+13)%26
		}

		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = 'A' + (p[i]-'A'+13)%26
		}
	}

	return n, err
}

func Rot13Run() {
	r := rot13Reader{
		r: strings.NewReader("Hello World!"),
	}

	io.Copy(os.Stdout, r)
}
