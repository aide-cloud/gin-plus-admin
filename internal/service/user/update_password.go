package user

import (
	"context"
)

type (
	// UpdatePasswordService ...
	UpdatePasswordParams struct {
		// add params
		ID          uint   `uri:"id"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	// UpdatePasswordReply ...
	UpdatePasswordReply struct {
		// add reply
	}
)

// UpdatePassword ...
func (l *User) UpdatePassword(ctx context.Context, req *UpdatePasswordParams) (*UpdatePasswordReply, error) {
	// add your code here
	return &UpdatePasswordReply{}, nil
}
