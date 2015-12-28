package main
import (
	"testing"
	"os"
	"io/ioutil"
	"strings"
)

func TestAppendToFile(t *testing.T){
	filename := "/tmp/test"
	info := "info"
	f, err := os.Create(filename)
	if err != nil {
		t.Error(err)
	}
	f.Close()

	appendToFile(&info, &filename)

	fb, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}

	if fb == nil {
		t.Error("A file must be returned")
	}

	fs := string(fb)

	if isInFile := strings.Contains(fs, "info"); isInFile == false {
		t.Error("String was not contained in file")
	}
}

func TestMountString(t *testing.T){
	args := []string{"abd","/123/"}

	s := mountString(args)

	expectedR := "export ABD=/123/\n"
	if *s !=  expectedR {
		t.Errorf("return \"%s\" != expected \"%s\"\n", *s, expectedR)
	}
}

func TestOpenFile(t *testing.T){
	fn := "/tmp/test"
	f, err := os.Create(fn)
	if err != nil {
		t.Error("File could not be created")
	}

	f.Close()

	f = openFile(fn)

	if f == nil {
		t.Error("A file must be returned")
	}

	//Now delete the file and pass it again, must return an error
	err = os.Remove(fn)
	if err != nil {
		t.Error("Error in test script: The file could not be removed")
	}

	defer func(){
		if err := recover(); err == nil {
			t.Error("No error has been returned after passing a non existing file")
		}
	}()

	f = openFile(fn)
}