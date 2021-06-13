package models

import (
	"context"
	"time"
)

type Base struct {
	Id int64 `json:"id" pg:"id"`
	CreatedAt time.Time  `json:"created_at" pg:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" pg:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" pg:"deleted_at"`
}


func (b *Base) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = now
	}
	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = now
	}

	return ctx, nil
}

func (b *Base) BeforeUpdate(ctx context.Context) (context.Context, error) {
	b.UpdatedAt = time.Now()
	return ctx, nil
}


func (b *Base) Delete() {
	t := time.Now()
	b.DeletedAt = &t
}