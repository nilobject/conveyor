package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestDestroy(t *testing.T) {
	err := migrateDatabaseDown()
	assert.NoError(t, err, "Error migrating down")
	err = migrateDatabaseUp()
	assert.NoError(t, err, "Error migrating up")
}

func TestDownMigrationWithBadInfo(t *testing.T) {
	config().pg.port = 1
	err := migrateDatabaseDown()
	assert.Error(t, err, "No error with bad connection string")
	initializeConfiguration()
}

func TestUpMigrationWithBadInfo(t *testing.T) {
	config().pg.port = 1
	err := migrateDatabaseUp()
	assert.Error(t, err, "No error with bad connection string")
	initializeConfiguration()
}

func TestConnection(t *testing.T) {
	db, _ := db()
	err := db.Ping()
	assert.NoError(t, err, "Error connecting to database")
}

func TestConnectionWithBadInfo(t *testing.T) {
	config().pg.port = 1
	db, _ := db()
	err := db.Ping()
	assert.Error(t, err, "No error with bad connection string")
	initializeConfiguration()
}
