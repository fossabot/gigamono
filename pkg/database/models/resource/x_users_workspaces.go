package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// XUsersWorkspaces is represents users membership to a workspace.
type XUsersWorkspaces struct {
	models.BaseNoID
	UserID      uuid.UUID `pg:",pk, type:uuid" json:"user_id"`
	WorkspaceID uuid.UUID `pg:",pk, type:uuid" json:"workspace_id"`
}
