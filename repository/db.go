package repository

import (
	sq "github.com/Masterminds/squirrel"
)

var (
	Psq = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)
