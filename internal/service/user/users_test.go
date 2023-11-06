package user_test

// import (
// 	"context"
// 	"testing"
// 	"time"
//
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/vladjong/go_project_template/internal/entity"
// 	"github.com/vladjong/go_project_template/internal/entity/dto"
// 	"github.com/vladjong/go_project_template/internal/repository/mocks"
// 	users_repository "github.com/vladjong/go_project_template/internal/repository/postgres/users"
// 	"github.com/vladjong/go_project_template/internal/services/users"
// )
//
// func Test_service_Users(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	type fields struct {
// 		repository *mocks.Repository
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		prepare func(f *fields)
// 		want    []entity.User
// 		wantErr bool
// 		err     error
// 	}{
// 		{
// 			name: "Ok",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			prepare: func(f *fields) {
// 				f.repository.
// 					On("Users", mock.Anything).
// 					Return([]dto.User{
// 						{Nickname: "test1", Birthday: time.Date(1999, time.January, 10, 0, 0, 0, 0, time.UTC)},
// 						{Nickname: "test2", Birthday: time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC)},
// 						{Nickname: "test3", Birthday: time.Date(2005, time.January, 11, 0, 0, 0, 0, time.UTC)},
// 					}, nil)
// 			},
// 			want: []entity.User{
// 				{Nickname: "test1", Birthday: time.Date(1999, time.January, 10, 0, 0, 0, 0, time.UTC)},
// 				{Nickname: "test2", Birthday: time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC)},
// 				{Nickname: "test3", Birthday: time.Date(2005, time.January, 11, 0, 0, 0, 0, time.UTC)},
// 			},
// 		},
// 		{
// 			name: "Error Not Found",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			prepare: func(f *fields) {
// 				f.repository.
// 					On("Users", mock.Anything).
// 					Return(nil, users_repository.ErrNotFound)
// 			},
// 			wantErr: true,
// 			err:     users_repository.ErrNotFound,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			f := fields{
// 				repository: mocks.NewRepository(t),
// 			}
// 			tt.prepare(&f)
//
// 			s := users.New(f.repository)
// 			actual, err := s.Users(tt.args.ctx)
// 			if (err != nil) == tt.wantErr {
// 				assert.ErrorIs(t, err, tt.err)
// 				return
// 			}
// 			assert.Equal(t, tt.want, actual)
// 		})
// 	}
// }
