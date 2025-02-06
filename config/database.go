package config

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DB struct {
	db         neo4j.DriverWithContext
	dbUri      string
	dbUser     string
	dbPassword string
}

func NewDB(dbUri, dbUser, dbPassword string) *DB {
	return &DB{
		dbUri:      dbUri,
		dbUser:     dbUser,
		dbPassword: dbPassword,
	}
}

func (db *DB) GetDB() neo4j.DriverWithContext {
	return db.db
}

func (db *DB) Connect(ctx context.Context) error {
	var err error
	db.db, err = neo4j.NewDriverWithContext(
		db.dbUri,
		neo4j.BasicAuth(db.dbUser, db.dbPassword, ""))
	if err != nil {
		return err
	}

	err = db.db.VerifyConnectivity(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Close(ctx context.Context) error {
	return db.db.Close(ctx)
}
