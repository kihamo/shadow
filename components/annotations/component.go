package annotations

import (
	"github.com/kihamo/shadow"
)

type Component interface {
	shadow.Component
	Storage
}
