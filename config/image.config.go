package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/apex/log"
)

type Manifest struct {
	MediaType   string            `json:"mediaType"`
	Digest      string            `json:"digest"`
	Size        int               `json:"size"`
	Annotations map[string]string `json:"annotations"`
}

type Index struct {
	SchemaVersion int        `json:"schemaVersion"`
	Manifests     []Manifest `json:"manifests"`
}

func GetImageTag(imagePath string) []string {
	jsonFile, err := os.Open(filepath.Join(imagePath, "index.json")) // index.json 파일 경로를 적절하게 수정하세요.
	if err != nil {
		log.Fatal(err.Error())
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var index Index
	json.Unmarshal(byteValue, &index)

	var tags []string
	for _, manifest := range index.Manifests {
		tags = append(tags, manifest.Annotations["org.opencontainers.image.ref.name"])
	}

	return tags
}

func RemoveImage(imageDir string, tag string) {
	indexFilePath := filepath.Join(imageDir, "index.json")

	file, err := os.Open(indexFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())

	}

	var index Index
	json.Unmarshal(byteValue, &index)

	var newManifests []Manifest
	for _, manifest := range index.Manifests {
		if manifest.Annotations["org.opencontainers.image.ref.name"] != tag {
			newManifests = append(newManifests, manifest)
		} else {
			blobPath := filepath.Join(imageDir, "blobs", "sha256", manifest.Digest[7:])
			err := os.Remove(blobPath)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
		}
	}
	index.Manifests = newManifests

	file, err = os.Create(indexFilePath)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(index)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
