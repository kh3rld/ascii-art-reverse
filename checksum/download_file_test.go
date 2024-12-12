package checksum

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	if err := os.MkdirAll("banners", 0755); err != nil {
		t.Fatalf("Failed to create banners directory: %v", err)
	}

	// Create a test server to simulate successful and failed downloads
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/banners/standard.txt":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Standard ASCII Art"))
		case "/banners/shadow.txt":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Shadow ASCII Art"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	// Override the URLs in the fileURLs map to point to the test server
	fileURLs["banners/standard.txt"] = ts.URL + "/banners/standard.txt"
	fileURLs["banners/shadow.txt"] = ts.URL + "/banners/shadow.txt"

	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Valid file - standard",
			args:    args{file: "banners/standard.txt"},
			wantErr: false,
		},
		{
			name:    "Valid file - shadow",
			args:    args{file: "banners/shadow.txt"},
			wantErr: false,
		},
		{
			name:    "Invalid file",
			args:    args{file: "banners/invalid.txt"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DownloadFile(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	for _, tt := range tests {
		if !tt.wantErr {
			_ = os.Remove(tt.args.file)
		}
	}
}
