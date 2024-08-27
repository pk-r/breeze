package action

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pk-r/breeze/pkg/database"
	"github.com/pk-r/breeze/pkg/storage"
	"gopkg.in/yaml.v3"
)

type Sync struct {
	Storage       storage.Storage
	JobRepository database.JobRepository
}

func (s Sync) Run(ctx context.Context) error {
	files, err := s.Storage.FetchFiles(ctx)

	if err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}

	for _, file := range files {
		var content map[string]interface{}

		err := yaml.Unmarshal(file, &content)
		if err != nil {
			return fmt.Errorf("failed to decode YAML: %w", err)
		}

		jobs := []database.Job{}
		invalidJobsCount := 0

		for title, value := range content {
			if job, ok := value.(map[string]interface{}); ok {
				if _, hasScript := job["script"]; hasScript {
					jobStruct, err := mapToStruct[database.Job](job)
					if err != nil {
						fmt.Println(fmt.Errorf("invalid job: %s, %w", title, err))
						invalidJobsCount++
						continue
					}

					jobStruct.Title = title
					jobs = append(jobs, jobStruct)
				}
			}
		}

		if len(jobs) == 0 {
			return nil
		}

		if err = s.JobRepository.Sync(jobs); err != nil {
			return fmt.Errorf("faield to sync jobs to database: %w", err)
		}
	}

	return nil
}

func mapToStruct[T any](data map[string]interface{}) (T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return *new(T), fmt.Errorf("error marshaling map: %w", err)
	}

	var result T
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return *new(T), fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return result, nil
}
