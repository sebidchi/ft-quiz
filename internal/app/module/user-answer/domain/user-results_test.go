package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersResults_UserScore(t *testing.T) {
	tests := []struct {
		name         string
		usersResults UsersResults
		userId       string
		want         *UserResults
		wantedError  error
	}{
		{
			name: "User with results",
			usersResults: UsersResults{
				"user1": Result{
					Points:         9,
					TotalQuestions: 10,
				},
				"user2": Result{
					Points:         7,
					TotalQuestions: 10,
				},
				"user3": Result{
					Points:         8,
					TotalQuestions: 10,
				},
			},
			userId: "user1",
			want: &UserResults{
				UserId:     "user1",
				Total:      9,
				Percentage: 90,
				BetterThan: 100,
				TotalUsers: 3,
			},
			wantedError: nil,
		},
		{
			name: "Single User results",
			usersResults: UsersResults{
				"user1": Result{
					Points:         9,
					TotalQuestions: 10,
				},
			},
			userId: "user1",
			want: &UserResults{
				UserId:     "user1",
				Total:      9,
				Percentage: 90,
				BetterThan: 0,
				TotalUsers: 1,
			},
			wantedError: nil,
		},
		{
			name: "User without results",
			usersResults: UsersResults{
				"user1": Result{
					Points:         8,
					TotalQuestions: 10,
				},
				"user2": Result{
					Points:         7,
					TotalQuestions: 10,
				},
				"user3": Result{
					Points:         9,
					TotalQuestions: 10,
				},
			},
			userId:      "user4",
			want:        nil,
			wantedError: NewUsersResultsNotFound("user4"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.usersResults.UserResults(tt.userId)

			if tt.wantedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
