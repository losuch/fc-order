package api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/losuch/fc-order/db/mock"
	db "github.com/losuch/fc-order/db/sqlc"
	"github.com/losuch/fc-order/pb"
	"github.com/losuch/fc-order/util"
	"github.com/stretchr/testify/require"
)

// TestCreateAccountAPI tests the CreateAccount function
func TestCreateAccountAPI(t *testing.T) {

	// create a random account
	account, _ := randomAccount(t)
	
	testCases := []struct {
		name        string
		request     *pb.CreateAccountRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, resp *pb.Account, err error)
	}{
		{
			name: "OK",
			request: &pb.CreateAccountRequest{
				Email:    account.Email,
				Password: account.HashedPassword,
				Role:     account.Role,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateAccountParams{
					Email:          account.Email,
					HashedPassword: account.HashedPassword,
					Role:           account.Role,
				}
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Account{
						AccountID:      account.AccountID,
						Email:          account.Email,
						HashedPassword: account.HashedPassword,
						Role:           account.Role,
					}, nil)				
			},
			checkResponse: func(t *testing.T, res *pb.Account, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotZero(t, res.AccountId)
				require.Equal(t, account.Email, res.Email)
				require.Equal(t, account.Role, res.Role)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()
			store := mockdb.NewMockStore(storeCtrl)

			taskCtrl := gomock.NewController(t)
			defer taskCtrl.Finish()

			tc.buildStubs(store)
			server := newTestServer(t, store)

			res, err := server.CreateAccount(context.Background(), tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}


func randomAccount(t *testing.T) (account db.Account, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)

	if err != nil {
		require.NoError(t, err)
	}

	account = db.Account{
		AccountID:      util.RandomInt64(1, 1000),
		Email:          util.RandomEmail(),
		HashedPassword: hashedPassword,
		Role:           util.RandomString(4),
	}
	return
}

// func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
// 	data, err := ioutil.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotAccount db.Account
// 	err = json.Unmarshal(data, &gotAccount)
// 	require.NoError(t, err)
// 	require.Equal(t, account, gotAccount)
// }
