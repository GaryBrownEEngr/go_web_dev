package utils

import (
	"fmt"
	"os"

	"github.com/GaryBrownEEngr/twertle_api_dev/backend/models"
	"github.com/GaryBrownEEngr/twertle_api_dev/backend/utils/stacktrs"
	"github.com/jszwec/csvutil"
)

func GitBuildHashGet() (models.GitBuildHash, error) {
	// Check all the way to 3 levels up for the file.
	data, err := os.ReadFile("git_hash.txt")
	if err != nil {
		data, err = os.ReadFile("../git_hash.txt")
		if err != nil {
			data, err = os.ReadFile("../../git_hash.txt")
			if err != nil {
				data, err = os.ReadFile("../../../git_hash.txt")
				if err != nil {
					return models.GitBuildHash{}, fmt.Errorf("git_hash.txt not found in ")
				}
			}
		}
	}

	var gitHash []models.GitBuildHash
	err = csvutil.Unmarshal(data, &gitHash)
	if err != nil {
		return models.GitBuildHash{}, stacktrs.Wrap(err)
	}

	if len(gitHash) != 1 {
		return models.GitBuildHash{}, stacktrs.Errorf("Wrong number of lines found in git_hash.txt, %d", len(gitHash))
	}

	return gitHash[0], nil
}
