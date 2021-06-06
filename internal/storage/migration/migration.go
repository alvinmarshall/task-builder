package migration

import (
	"context"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx := db.WithContext(ctx)
	if err := tx.AutoMigrate(models...); err != nil {
		return err
	}
	cancel()
	return nil
}
