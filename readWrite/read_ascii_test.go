package readwrite

import (
	"os"
	"reflect"
	"testing"
)

// TestReadAscii tests the ReadAscii function for various scenarios.
func TestReadAscii(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     [][]string
		wantErr  bool
	}{
		{
			name:     "Invalid file name",
			filename: "nonexistent.txt",
			wantErr:  true,
		},
		{
			name:     "Unsupported file format",
			filename: "shadow.jpg",
			wantErr:  true,
		},
		{
			name:     "Unsupported file name in map",
			filename: "invalidfile.txt",
			wantErr:  true,
		},
		{
			name:     "File read error",
			filename: "readonly.txt",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Valid file" {
				tmpfile, err := os.CreateTemp("", "testfile*.txt")
				if err != nil {
					t.Fatal(err)
				}
				defer os.Remove(tmpfile.Name())

				data := ` _
| | 
| | 
| | 
|_| 
(_)`
				_, err = tmpfile.WriteString(data)
				if err != nil {
					t.Fatal(err)
				}

				got, err := ReadAscii(tmpfile.Name())
				if (err != nil) != tt.wantErr {
					t.Errorf("ReadAscii() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReadAscii() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
