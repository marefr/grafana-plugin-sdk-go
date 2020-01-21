package adapter

import (
	"context"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/backend/models"

	"github.com/grafana/grafana-plugin-sdk-go/genproto/pluginv2"
	"github.com/stretchr/testify/require"
)

func TestGetSchema(t *testing.T) {
	t.Run("GetSchema without a schema provider should return empty schema", func(t *testing.T) {
		adapter := &SDKAdapter{}
		res, err := adapter.GetSchema(context.Background(), &pluginv2.GetSchema_Request{})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Len(t, res.Resources, 0)
	})

	t.Run("GetSchema with a schema provider should return schema", func(t *testing.T) {
		adapter := &SDKAdapter{
			SchemaProvider: func() models.Schema {
				return models.Schema{
					Resources: models.ResourceMap{
						"test": models.NewResource("/"),
					},
				}
			},
		}
		res, err := adapter.GetSchema(context.Background(), &pluginv2.GetSchema_Request{})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Len(t, res.Resources, 1)
		require.Equal(t, "/", res.Resources["test"].Path)
	})
}
