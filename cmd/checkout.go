package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/chrillux/go-wyag/object"
	"github.com/spf13/cobra"
)

// catfileCmd represents the catfile command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout the contents of a commit object",
	Long:  `Checkout the contents of a commit object, the equivalent to git checkout.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		checkout(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func checkout(commitsha, path string) {
	obj, err := object.ReadObject(commitsha)
	if err != nil {
		log.Fatal(err)
	}

	if obj.GetObjType() != "commit" {
		log.Fatalf("object %s is not a commit object", commitsha)
	}

	commitObject := obj.(*object.Commit)
	treeObject, err := findTreeObject(commitObject)
	if err != nil {
		log.Fatal(err)
	}

	fileinfo, err := os.Stat(path)
	if err == nil { // file/dir exists
		if !fileinfo.IsDir() {
			log.Fatalf("path is not a directory: %s", path)
		}
		f, err := os.Open(path)
		if err != nil {
			log.Fatalf("could not open path: %s", path)
		}
		defer f.Close()

		// try to read one file from dir. If the dir is empty it will return a io.EOF error.
		_, err = f.Readdir(1)
		if err != io.EOF {
			log.Fatalf("%s is not an empty dir", path)
		}
	} else {
		// path does not exist, let's create that path
		os.MkdirAll(path, 0755)
	}

	treeCheckout(treeObject, path)
}

func findTreeObject(commitObject *object.Commit) (*object.Tree, error) {
	for _, kvlm := range commitObject.KVLM().KeyValues {
		if kvlm.Key == "tree" {
			to, err := object.ReadObject(kvlm.Value)
			if err != nil {
				return nil, fmt.Errorf("could not read tree object")
			}
			treeObject, ok := to.(*object.Tree)
			if !ok {
				return nil, fmt.Errorf("invalid tree object")
			}
			return treeObject, nil
		}
	}
	return nil, fmt.Errorf("could not find tree object")
}

func treeCheckout(treeObject *object.Tree, path string) {
	for _, item := range treeObject.Items() {
		o, err := object.ReadObject(item.SHA())
		if err != nil {
			log.Fatalf("could not read object %s", item.SHA())
		}
		dest := filepath.Join(path, item.Path())
		if o.GetObjType() == "tree" {
			treeobj := o.(*object.Tree)
			os.MkdirAll(dest, 0755)
			treeCheckout(treeobj, dest)
		} else if o.GetObjType() == "blob" {
			data, err := ioutil.ReadAll(o.Serialize())
			if err != nil {
				log.Fatalf("could not read file %s", dest)
			}
			err = os.WriteFile(dest, data, 0644)
			if err != nil {
				log.Fatalf("could not write file %s", dest)
			}
		}
	}
}
