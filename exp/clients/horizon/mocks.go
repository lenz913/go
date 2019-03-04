package horizonclient

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockClient is a mockable horizon client.
type MockClient struct {
	mock.Mock
}

// AccountDetail is a mocking method
func (m *MockClient) AccountDetail(request AccountRequest) (Account, error) {
	a := m.Called(request)
	return a.Get(0).(Account), a.Error(1)
}

// AccountData is a mocking method
func (m *MockClient) AccountData(request AccountRequest) (AccountData, error) {
	a := m.Called(request)
	return a.Get(0).(AccountData), a.Error(1)
}

// Effects is a mocking method
func (m *MockClient) Effects(request EffectRequest) (EffectsPage, error) {
	a := m.Called(request)
	return a.Get(0).(EffectsPage), a.Error(1)
}

// Assets is a mocking method
func (m *MockClient) Assets(request AssetRequest) (AssetsPage, error) {
	a := m.Called(request)
	return a.Get(0).(AssetsPage), a.Error(1)
}

func (m *MockClient) Stream(
	request StreamRequest, ctx context.Context,
	handler func(interface{}),
) error {
	a := m.Called(request, ctx, handler)
	return a.Error(0)
}

// ensure that the MockClient implements ClientInterface
var _ ClientInterface = &MockClient{}
