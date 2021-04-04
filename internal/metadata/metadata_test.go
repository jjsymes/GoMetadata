package metadata

import (
	"testing"
)

// func TestGetMetadata_BadData(t *testing.T) {
// 	cases := []struct {
// 		Name  string
// 		Error string
// 	}{
// 		{"../../test/fixtures/invalidformat.txt", "not a valid file format"},
// 	}
// 	for i, c := range cases {
// 		f, err := os.Open(c.Name)
// 		if err != nil {
// 			t.Fatalf("err: %v", err)
// 		}
// 		fmt.Println(i)
// 		defer f.Close()
// 		// if err != nil || !strings.Contains(err.Error(), c.Error) {
// 		// 	t.Fatalf("case %d (%s): %v", i, c.Name, err)
// 		// }
// 	}
// }

func TestFileExists_DoesNotExist(t *testing.T) {
	output := fileExists("dsfaf")
	if output != false {
		t.Fatalf("Expected false but got: %v", output)
	}
}

func TestFileExists_DoesExist(t *testing.T) {
	output := fileExists("../../test/fixtures/exists.txt")
	if output != true {
		t.Fatalf("Expected true but got: %v", output)
	}
}

func TestValidFile_DoesNotExist(t *testing.T) {
	output, invalidFileReason := validFile("../../test/fixtures/doesnotexistfsak")
	if (output != false || invalidFileReason != "File does not exist") {
		t.Fatalf("Expected false with file does not exist but got: %v %v", output, invalidFileReason)
	}
}

func TestValidFile_NotValid(t *testing.T) {
	output, invalidFileReason := validFile("../../test/fixtures/invalidFileExtension.txt")
	if (output != false || invalidFileReason[0:26] != "Not a valid file extension") {
		t.Fatalf("Expected false with Not a valid file extension but got: %v %v", output, invalidFileReason[0:26])
	}
}
