package multiwriter

import (
	"io"
	"testing"

	"dat515/2go/errors"

	"github.com/google/go-cmp/cmp"
)

func testMultiWriter(t *testing.T, scoreDec func()) {
	for _, test := range multiWriterTests {
		n, errs := WriteTo(test.inb, test.inw...)
		if !cmp.Equal(n, test.wantN) {
			if n == nil {
				t.Errorf("WriteTo(%q, writers...) = nil, want %v", test.inb, test.wantN)
			} else {
				t.Errorf("WriteTo(%q, writers...) = %v, want %v", test.inb, n, test.wantN)
			}
			scoreDec()
			continue
		}
		if diff := cmp.Diff(test.wantErr, errs, cmp.Comparer(errorsComparer)); diff != "" {
			t.Errorf("Error(): (-want +got):\n%s", diff)
			scoreDec()
		}
	}
}

var multiWriterTests = []struct {
	inb     []byte
	inw     []io.Writer
	wantN   []int
	wantErr errors.Errors
}{
	{nil, nil, []int{}, nil},
	{nil, []io.Writer{}, []int{}, nil},
	{nil, []io.Writer{io.Discard}, []int{0}, nil},
	{[]byte(""), []io.Writer{io.Discard}, []int{0}, nil},
	{[]byte("\n"), []io.Writer{io.Discard}, []int{1}, nil},
	{[]byte("TEST-001\n"), []io.Writer{io.Discard}, []int{9}, nil},
	{[]byte("TEST-002\n"), []io.Writer{io.Discard}, []int{9}, nil},
	{[]byte("TEST-003\n"), []io.Writer{io.Discard, io.Discard}, []int{9, 9}, nil},
	{[]byte("TEST-004\n"), []io.Writer{io.Discard, io.Discard, io.Discard}, []int{9, 9, 9}, nil},
	{
		[]byte("TEST-005\n"),
		[]io.Writer{
			io.Discard,
			failureWriter(2),
			io.Discard,
		},
		[]int{9, 2, 9},
		[]error{nil, io.ErrShortWrite, nil},
	},
	{
		[]byte("TEST-006\n"),
		[]io.Writer{
			io.Discard,
			failureWriter(2),
			failureWriter(6),
		},
		[]int{9, 2, 6},
		[]error{nil, io.ErrShortWrite, io.ErrShortWrite},
	},
	{
		[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		[]io.Writer{
			io.Discard,
			failureWriter(2),
			failureWriter(20),
		},
		[]int{26, 2, 20},
		[]error{nil, io.ErrShortWrite, io.ErrShortWrite},
	},
}
