package cli

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const VERSION = "v0.2.1"

func (z *ZVM) fetchVersionMap() (zigVersionMap, error) {

	req, err := http.NewRequest("GET", "https://ziglang.org/download/index.json", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "zvm (Zig Version Manager) " + VERSION)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	versions, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(filepath.Join(z.zvmBaseDir, "versions.json"), versions, 0755); err != nil {
		return nil, err
	}

	rawVersionStructure := make(zigVersionMap)
	if err := json.Unmarshal(versions, &rawVersionStructure); err != nil {
		return nil, err
	}

	return rawVersionStructure, nil
}
