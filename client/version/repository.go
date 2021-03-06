package version

import (
	"bufio"
	"github.com/kowala-tech/kcoin/client/log"
	"net/http"
)

type AssetRepository interface {
	All() ([]Asset, error)
}

const indexFile = "index.txt"

type s3assetRepository struct {
	repository string
}

func NewS3AssetRepository(path string) AssetRepository {
	return s3assetRepository{
		repository: path,
	}
}

func (ar s3assetRepository) All() ([]Asset, error) {
	response, err := http.Get(ar.repository + "/" + indexFile)
	if err != nil {
		return []Asset{}, err
	}
	defer response.Body.Close()

	var assets []Asset
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		filename := scanner.Text()
		version, err := FilenameParser(filename)
		if err != nil {
			// ignore error and continue to next filename
			log.Debug(
				"file is not a binary asset",
				"filename", filename,
				"err", err)
			continue
		}
		assets = append(assets, version)
	}

	if err := scanner.Err(); err != nil {
		return []Asset{}, err
	}

	return assets, nil
}

type memoryAssetRepository struct {
	assets []Asset
}

func NewMemoryAssetRepository(assets []Asset) AssetRepository {
	return memoryAssetRepository{
		assets: assets,
	}
}

func (r memoryAssetRepository) All() ([]Asset, error) {
	return r.assets, nil
}
