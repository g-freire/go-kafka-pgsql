package postgres

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
)

func DoMigrate(migrationsRootFolder, databaseURL string, log logger.Logger) error {
	m, err := migrate.New(
		migrationsRootFolder,
		databaseURL,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Errorf("MIGRATION ERROR: \n", err)
		return err
	}
	log.Debug("MIGRATION WORKS")
	return nil
}
