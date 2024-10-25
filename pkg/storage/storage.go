// pkg/storage/storage.go

package storage

import (
    "encoding/json"
    "os"
    "io/ioutil" // - obsolete? not available?
    // "errors"
)

func SaveData(filePath string, data interface{}) error {
    file, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filePath, file, 0644)
}

func LoadData(filePath string, v interface{}) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }

    return json.Unmarshal(byteValue, v)
}
