package adapter

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/genproto/pluginv2"
)

func TestCallResource(t *testing.T) {
	t.Run("Test call resource basic", func(t *testing.T) {
		anyHandler := &TestResourceHandler{}
		getHandler := &TestResourceHandler{}
		putHandler := &TestResourceHandler{}
		postHandler := &TestResourceHandler{}
		deleteHandler := &TestResourceHandler{}
		patchHandler := &TestResourceHandler{}
		adapter := &SDKAdapter{
			schema: backend.Schema{
				Resources: backend.ResourceMap{
					"test": backend.NewResource("/").
						AddRoute("/", backend.RouteMethodAny, anyHandler.handle).
						AddRoute("/", backend.RouteMethodGet, getHandler.handle).
						AddRoute("/", backend.RouteMethodPut, putHandler.handle).
						AddRoute("/", backend.RouteMethodPost, postHandler.handle).
						AddRoute("/", backend.RouteMethodDelete, deleteHandler.handle).
						AddRoute("/", backend.RouteMethodPatch, patchHandler.handle),
				},
			},
		}

		t.Run("Call non-registered resource should return 404", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "non-existing",
			})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusNotFound, int(res.Code))
			assert.Equal(t, 0, anyHandler.callerCount)
		})

		t.Run("Call test resource should call any handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/",
				Method:       http.MethodTrace,
			})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, int(res.Code))
			assert.Equal(t, 1, anyHandler.callerCount)
		})

		t.Run("Call test resource should call get handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/",
				Method:       http.MethodGet,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, getHandler.callerCount)
		})

		t.Run("Call test resource should call put handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/",
				Method:       http.MethodPut,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, putHandler.callerCount)
		})

		t.Run("Call test resource should call post handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/",
				Method:       http.MethodPost,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, postHandler.callerCount)
		})

		t.Run("Call test resource should call delete handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/",
				Method:       http.MethodDelete,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, deleteHandler.callerCount)
		})

		t.Run("Call test resource should call patch handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/",
				Method:       http.MethodPatch,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, patchHandler.callerCount)
		})
	})

	t.Run("Test call resource advanced", func(t *testing.T) {
		anyHandler := &TestResourceHandler{}
		getHandler := &TestResourceHandler{}
		putHandler := &TestResourceHandler{}
		postHandler := &TestResourceHandler{}
		deleteHandler := &TestResourceHandler{}
		patchHandler := &TestResourceHandler{}
		adapter := &SDKAdapter{
			schema: backend.Schema{
				Resources: backend.ResourceMap{
					"test": backend.NewResource("/test/:id").
						AddRoute("/", backend.RouteMethodAny, anyHandler.handle).
						AddRoute("/get", backend.RouteMethodGet, getHandler.handle).
						AddRoute("/put", backend.RouteMethodPut, putHandler.handle).
						AddRoute("/post", backend.RouteMethodPost, postHandler.handle).
						AddRoute("/delete", backend.RouteMethodDelete, deleteHandler.handle).
						AddRoute("/patch", backend.RouteMethodPatch, patchHandler.handle),
				},
			},
		}

		t.Run("Call non-registered resource should return 404", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "non-existing",
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(404), res.Code)
			assert.Equal(t, 0, anyHandler.callerCount)
		})

		t.Run("Call test resource should call any handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/test/:id",
				Method:       http.MethodTrace,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, anyHandler.callerCount)
		})

		t.Run("Call test resource should call get handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/test/:id/get",
				Method:       http.MethodGet,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, getHandler.callerCount)
		})

		t.Run("Call test resource should call put handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/test/:id/put",
				Method:       http.MethodPut,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, putHandler.callerCount)
		})

		t.Run("Call test resource should call post handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/test/:id/post",
				Method:       http.MethodPost,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, postHandler.callerCount)
		})

		t.Run("Call test resource should call delete handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/test/:id/delete",
				Method:       http.MethodDelete,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, deleteHandler.callerCount)
		})

		t.Run("Call test resource should call patch handler and return 200", func(t *testing.T) {
			res, err := adapter.CallResource(context.Background(), &pluginv2.CallResource_Request{
				Config:       &pluginv2.PluginConfig{},
				ResourceName: "test",
				ResourcePath: "/test/:id/patch",
				Method:       http.MethodPatch,
			})
			assert.NoError(t, err)
			assert.Equal(t, int32(200), res.Code)
			assert.Equal(t, 1, patchHandler.callerCount)
		})
	})
}

type TestResourceHandler struct {
	callerCount int
}

func (h *TestResourceHandler) handle(resourceCtx *backend.ResourceRequestContext) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		h.callerCount++
		rw.WriteHeader(200)
	})
}
