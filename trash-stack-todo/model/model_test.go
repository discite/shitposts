package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	// Test case: successful setup
	db, err := Setup()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Test case: setup called multiple times, should return the same db instance
	db2, err := Setup()
	assert.NoError(t, err)
	assert.Equal(t, db, db2)
}

func TestMakeMigrations(t *testing.T) {
	// Test case: successful migrations
	err := MakeMigrations()
	assert.NoError(t, err)

	// Test case: table already exists, should not return an error
	err = MakeMigrations()
	assert.NoError(t, err)
}
