package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	conf "github.com/ipfs/ipfs-ds-convert/config"

	config "gx/ipfs/QmaZiW9UbccSq3zp1nhYorSpYBUjsPG8tyfTE1nxBFkNEx/go-ipfs/repo/config"
	fsrepo "gx/ipfs/QmaZiW9UbccSq3zp1nhYorSpYBUjsPG8tyfTE1nxBFkNEx/go-ipfs/repo/fsrepo"
)

func NewTestRepo(t *testing.T, spec map[string]interface{}) (string, func(t *testing.T)) {
	conf, err := config.Init(os.Stdout, 1024)
	if err != nil {
		t.Fatal(err)
	}

	err = config.ConfigProfiles["test"](conf)
	if err != nil {
		t.Fatal(err)
	}

	if spec != nil {
		conf.Datastore.Spec = spec
	}

	repoRoot, err := ioutil.TempDir("/tmp", "ds-convert-test-")
	if err != nil {
		t.Fatal(err)
	}

	if err := fsrepo.Init(repoRoot, conf); err != nil {
		t.Fatal(err)
	}

	return repoRoot, func(t *testing.T) {
		err := os.RemoveAll(repoRoot)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func PatchConfig(t *testing.T, configPath string, newSpecPath string) {
	newSpec := make(map[string]interface{})
	err := conf.Load(newSpecPath, &newSpec)
	if err != nil {
		t.Fatal(err)
	}

	repoConfig := make(map[string]interface{})
	err = conf.Load(configPath, &repoConfig)
	if err != nil {
		t.Fatal(err)
	}

	dsConfig, ok := repoConfig["Datastore"].(map[string]interface{})
	if !ok {
		t.Fatal(fmt.Errorf("no 'Datastore' or invalid type in %s", configPath))
	}

	_, ok = dsConfig["Spec"].(map[string]interface{})
	if !ok {
		t.Fatal(fmt.Errorf("no 'Datastore.Spec' or invalid type in %s", configPath))
	}

	dsConfig["Spec"] = newSpec

	b, err := json.MarshalIndent(repoConfig, "", "  ")
	ioutil.WriteFile(configPath, b, 0660)
}
