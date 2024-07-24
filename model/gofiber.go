package model

import (
	"github.com/gofiber/fiber/v2"
)

type GofiberMetadata struct {
	Context *fiber.Ctx
	Metadata
}
