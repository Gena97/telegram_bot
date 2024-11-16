package repository

import (
	"encoding/json"
	"fmt"

	"github.com/Gena97/telegram_bot/internal/service"
)

func (r *PGXRepository) GetFileHashes(lenHashes int, file_type string) ([]service.FileHash, error) {
	// Prepare a query that retrieves hash values and message IDs.
	rows, err := r.mainDB.Query(`SELECT hash_value, message_id FROM hashes WHERE hash_count = $1 AND file_type = $2`, lenHashes, file_type)
	if err != nil {
		return nil, fmt.Errorf("error querying hash from database: %w", err)
	}
	defer rows.Close()

	var result []service.FileHash

	// Iterate over each row to process hash values and message IDs.
	for rows.Next() {
		var hashJSON string
		var messageID int64

		// Scan hash JSON and message ID into local variables.
		if err := rows.Scan(&hashJSON, &messageID); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Unmarshal JSON array of hashes into a slice of uint64.
		var hashes []uint64
		if err := json.Unmarshal([]byte(hashJSON), &hashes); err != nil {
			return nil, fmt.Errorf("error unmarshaling JSON to hashes: %w", err)
		}

		// Append the struct to the result slice.
		result = append(result, service.FileHash{
			Hashes:    hashes,
			MessageID: messageID,
		})
	}

	// Check for any errors encountered during the iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return result, nil
}

func (r *PGXRepository) InsertHash(fileHash service.FileHash) error {
	if fileHash.MessageID <= 0 {
		return fmt.Errorf("не заполнено MessageID")
	}
	if len(fileHash.Hashes) == 0 {
		return fmt.Errorf("не заполнено Hashes")
	}
	if !(fileHash.Type == "photo" || fileHash.Type == "video") {
		return fmt.Errorf("неизвестный Type хеша")
	}

	hashesJSON, err := json.Marshal(fileHash.Hashes)
	if err != nil {
		return fmt.Errorf("error marshaling hashes to JSON: %w", err)
	}

	// Вставляем JSON-строку в базу данных
	_, err = r.mainDB.Exec(`INSERT INTO hashes (message_id, file_type, hash_value, hash_count) VALUES (?, ?, ?, ?)`,
		fileHash.MessageID, fileHash.Type, string(hashesJSON), len(fileHash.Hashes))
	return err
}
