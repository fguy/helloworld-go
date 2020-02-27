package helloworld

import (
	"context"

	"github.com/fguy/helloworld-go/entities"
)

//go:generate mockgen -destination=../../mocks/repositories/helloworld/interface.go github.com/fguy/helloworld-go/repositories/helloworld Interface

// Interface is a interface of helloworld repository
type Interface interface {
	GetPage(context.Context, string) (*entities.Page, error)
}
