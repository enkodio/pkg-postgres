package client

import "github.com/jackc/pgx/v5/pgconn"

type CommandTag struct {
	pgconn.CommandTag
}

func NewCommandTag(commandTag pgconn.CommandTag, err error) (CommandTag, error) {
	return CommandTag{
		CommandTag: commandTag,
	}, err
}
