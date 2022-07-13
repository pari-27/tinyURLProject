package store

import (
	"github.com/tinyURLProject/utils/types"
)

type store interface {
	Create(types.UrlMap) (err error)
	Get()
	Update()
	Delete()
}
