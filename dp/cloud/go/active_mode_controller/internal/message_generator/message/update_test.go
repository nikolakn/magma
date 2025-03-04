package message_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"magma/dp/cloud/go/active_mode_controller/internal/message_generator/message"
	"magma/dp/cloud/go/active_mode_controller/protos/active_mode"
)

func TestUpdateMessageString(t *testing.T) {
	m := message.NewUpdateMessage(id)
	expected := fmt.Sprintf("update: %d", id)
	assert.Equal(t, expected, m.String())
}

func TestUpdateMessageSend(t *testing.T) {
	client := &stubUpdateClient{}
	provider := &stubUpdateClientProvider{client: client}

	m := message.NewUpdateMessage(id)
	require.NoError(t, m.Send(context.Background(), provider))

	expected := &active_mode.AcknowledgeCbsdUpdateRequest{Id: id}
	assert.Equal(t, expected, client.req)
}

type stubUpdateClientProvider struct {
	message.ClientProvider
	client *stubUpdateClient
}

func (s *stubUpdateClientProvider) GetActiveModeClient() active_mode.ActiveModeControllerClient {
	return s.client
}

type stubUpdateClient struct {
	active_mode.ActiveModeControllerClient
	req *active_mode.AcknowledgeCbsdUpdateRequest
}

func (s *stubUpdateClient) AcknowledgeCbsdUpdate(_ context.Context, in *active_mode.AcknowledgeCbsdUpdateRequest, _ ...grpc.CallOption) (*empty.Empty, error) {
	s.req = in
	return &empty.Empty{}, nil
}
