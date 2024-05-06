package huggingfacehub

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	// Set up test parameters
	repoId := "prajjwal1/bert-small"
	repoType := "model"
	revision := "main"
	fileName := "pytorch_model.bin"
	cacheDir := filepath.Join(os.TempDir(), "huggingface_cache")
	token := ""
	forceDownload := false
	forceLocal := false

	// Create a new HTTP client
	client := &http.Client{}

	// Define a progress function
	progressFn := func(progress, downloaded, total int, eof bool) {
		// TODO
	}

	// Call the Download function
	filePath, commitHash, err := Download(context.Background(), client, repoId, repoType, revision, fileName, cacheDir, token, forceDownload, forceLocal, progressFn)
	if err != nil {
		t.Fatalf("Download failed: %v", err)
	}

	// Check if the file was downloaded successfully
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("Downloaded file does not exist: %s", filePath)
	}

	// Check if the commit hash is not empty
	if commitHash == "" {
		t.Error("Commit hash is empty")
	}

	// Clean up the downloaded file and cache directory
	defer func() {
		os.RemoveAll(cacheDir)
	}()
}
