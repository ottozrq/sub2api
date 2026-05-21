package service

import (
	"context"
	"testing"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/pkg/pagination"
	"github.com/stretchr/testify/require"
)

type dispositionUserRepoStub struct {
	user *User
}

func (s *dispositionUserRepoStub) Create(context.Context, *User) error { panic("unexpected") }
func (s *dispositionUserRepoStub) GetByID(context.Context, int64) (*User, error) {
	return s.user, nil
}
func (s *dispositionUserRepoStub) GetByEmail(context.Context, string) (*User, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) GetFirstAdmin(context.Context) (*User, error) { panic("unexpected") }
func (s *dispositionUserRepoStub) Update(context.Context, *User) error          { panic("unexpected") }
func (s *dispositionUserRepoStub) Delete(context.Context, int64) error          { panic("unexpected") }
func (s *dispositionUserRepoStub) GetUserAvatar(context.Context, int64) (*UserAvatar, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) UpsertUserAvatar(context.Context, int64, UpsertUserAvatarInput) (*UserAvatar, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) DeleteUserAvatar(context.Context, int64) error { panic("unexpected") }
func (s *dispositionUserRepoStub) List(context.Context, pagination.PaginationParams) ([]User, *pagination.PaginationResult, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) ListWithFilters(context.Context, pagination.PaginationParams, UserListFilters) ([]User, *pagination.PaginationResult, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) GetLatestUsedAtByUserIDs(context.Context, []int64) (map[int64]*time.Time, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) GetLatestUsedAtByUserID(context.Context, int64) (*time.Time, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) UpdateUserLastActiveAt(context.Context, int64, time.Time) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) UpdateBalance(context.Context, int64, float64) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) DeductBalance(context.Context, int64, float64) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) UpdateConcurrency(context.Context, int64, int) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) BatchSetConcurrency(context.Context, []int64, int) (int, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) BatchAddConcurrency(context.Context, []int64, int) (int, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) ExistsByEmail(context.Context, string) (bool, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) RemoveGroupFromAllowedGroups(context.Context, int64) (int64, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) AddGroupToAllowedGroups(context.Context, int64, int64) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) RemoveGroupFromUserAllowedGroups(context.Context, int64, int64) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) ListUserAuthIdentities(context.Context, int64) ([]UserAuthIdentityRecord, error) {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) UnbindUserAuthProvider(context.Context, int64, string) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) UpdateTotpSecret(context.Context, int64, *string) error {
	panic("unexpected")
}
func (s *dispositionUserRepoStub) EnableTotp(context.Context, int64) error  { panic("unexpected") }
func (s *dispositionUserRepoStub) DisableTotp(context.Context, int64) error { panic("unexpected") }

func TestApplyUserDispositionRejectsDestructiveAdminActions(t *testing.T) {
	repo := &dispositionUserRepoStub{user: &User{ID: 1, Role: RoleAdmin, Status: StatusActive}}
	svc := &adminServiceImpl{userRepo: repo}

	_, err := svc.ApplyUserDisposition(context.Background(), UserDispositionInput{
		UserID:         1,
		Reason:         "risk action",
		DisableAPIKeys: true,
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "cannot apply destructive disposition actions to admin users")
}
