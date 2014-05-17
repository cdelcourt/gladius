package chef

import (
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestNodeFromFile(t *testing.T) {
	if n1, err := NodeFromFile("test/node.json"); err != nil {
		t.Fatal(err)
	} else {
		spew.Dump(n1)
	}
}

func TestNodeWriteToFile(t *testing.T) {
	n1 := &Node{
		"name":     "foo",
		"run_list": []string{"base", "java"},
	}
	spew.Dump(n1)

	tf, _ := ioutil.TempFile("test", "gladius-chef-node")
	defer tf.Close()
	defer os.Remove(tf.Name())

	// var b = new(bytes.Buffer)
	// _, err := b.ReadFrom(n1)
	// spew.Dump("b is", b.String())
	// spew.Dump("err is", err)

	// // because Node has a io.Reader Read() compliant implementation, we can copy out of it
	// // This hangs -- why?
	// if w, err := io.Copy(tf, n1); err != nil {
	// 	t.Errorf("could not copy node into tempfile, err: %v, written: %v\n", err, w)
	// } else {
	// 	spew.Dump(w)
	// 	spew.Dump(err)
	// }

	// if node, err := NodeFromFile(tf.Name()); err != nil {
	// 	t.Error("could not reserialize node from tempfile after writing it", err)
	// } else {
	// 	spew.Dump(node)
	// }

}
