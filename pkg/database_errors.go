package pkg

import "errors"

var ErrQueryExecution = errors.New("database: unable to run a query")
