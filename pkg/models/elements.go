// pkg/models/elements.go

package models

type Element struct {
    ID           string   `json:"id"`
    Type         string   `json:"type"`
    Name         string   `json:"name"`
    Descr        string   `json:"desc`
    Tags         []string `json:"tags"`
    TmpTags      []string `json:"tmp_tags"`
    Dependencies []string `json:"dependencies"`
}

type Credentials struct {
    ID           string   `json:"id"`
    Type         string   `json:"type"`
    Name         string   `json:"name"`
    Descr        string   `json:"desc`
    Tags         []string `json:"tags"`
    TmpTags      []string `json:"tmp_tags"`
    UsedIn       []string `json:"used_in"`
}

type Cluster struct {
    ID           string   `json:"id"`
    Type         string   `json:"type"`
    Name         string   `json:"name"`
    Descr        string   `json:"desc`
    Tags         []string `json:"tags"`
    TmpTags      []string `json:"tmp_tags"`
    Dependencies []string `json:"dependencies"`
}

type Node struct {
    ID           string   `json:"id"`
    Type         string   `json:"type"`
    Name         string   `json:"name"`
    Descr        string   `json:"desc`
    Tags         []string `json:"tags"`
    TmpTags      []string `json:"tmp_tags"`
    Dependencies []string `json:"dependencies"`
}

type Server struct {
    ID          string   `json:"id"`
    Type        string   `json:"type"`
    Name        string   `json:"name"`
    IPAddress   string   `json:"ip_address"`
    OS          string   `json:"operating_system"`
    Dependencies []string `json:"dependencies"`
}

type Application struct {
    ID          string   `json:"id"`
    Type        string   `json:"type"`
    Name        string   `json:"name"`
    Version     string   `json:"version"`
    Dependencies []string `json:"dependencies"`
}

type Library struct {
    ID          string   `json:"id"`
    Type        string   `json:"type"`
    Name        string   `json:"name"`
    Version     string   `json:"version"`
    Dependencies []string `json:"dependencies"`
}
