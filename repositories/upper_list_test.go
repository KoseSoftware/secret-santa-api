package repositories_test

import "github.com/KoseSoftware/secret-santa-api/repositories"

// test list repository implements lister interface
var _ repositories.ListerRepository = (*repositories.UpperListRepository)(nil)