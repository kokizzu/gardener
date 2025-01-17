// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/kubernetes (interfaces: ChartApplier)
//
// Generated by this command:
//
//	mockgen -package mock -destination=mocks_chartapplier.go github.com/gardener/gardener/pkg/client/kubernetes ChartApplier
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	embed "embed"
	reflect "reflect"

	chartrenderer "github.com/gardener/gardener/pkg/chartrenderer"
	kubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	gomock "go.uber.org/mock/gomock"
)

// MockChartApplier is a mock of ChartApplier interface.
type MockChartApplier struct {
	ctrl     *gomock.Controller
	recorder *MockChartApplierMockRecorder
	isgomock struct{}
}

// MockChartApplierMockRecorder is the mock recorder for MockChartApplier.
type MockChartApplierMockRecorder struct {
	mock *MockChartApplier
}

// NewMockChartApplier creates a new mock instance.
func NewMockChartApplier(ctrl *gomock.Controller) *MockChartApplier {
	mock := &MockChartApplier{ctrl: ctrl}
	mock.recorder = &MockChartApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChartApplier) EXPECT() *MockChartApplierMockRecorder {
	return m.recorder
}

// ApplyFromArchive mocks base method.
func (m *MockChartApplier) ApplyFromArchive(ctx context.Context, archive []byte, namespace, name string, opts ...kubernetes.ApplyOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, archive, namespace, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyFromArchive", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyFromArchive indicates an expected call of ApplyFromArchive.
func (mr *MockChartApplierMockRecorder) ApplyFromArchive(ctx, archive, namespace, name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, archive, namespace, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyFromArchive", reflect.TypeOf((*MockChartApplier)(nil).ApplyFromArchive), varargs...)
}

// ApplyFromEmbeddedFS mocks base method.
func (m *MockChartApplier) ApplyFromEmbeddedFS(ctx context.Context, embeddedFS embed.FS, chartPath, namespace, name string, opts ...kubernetes.ApplyOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, embeddedFS, chartPath, namespace, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyFromEmbeddedFS", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyFromEmbeddedFS indicates an expected call of ApplyFromEmbeddedFS.
func (mr *MockChartApplierMockRecorder) ApplyFromEmbeddedFS(ctx, embeddedFS, chartPath, namespace, name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, embeddedFS, chartPath, namespace, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyFromEmbeddedFS", reflect.TypeOf((*MockChartApplier)(nil).ApplyFromEmbeddedFS), varargs...)
}

// DeleteFromArchive mocks base method.
func (m *MockChartApplier) DeleteFromArchive(ctx context.Context, archive []byte, namespace, name string, opts ...kubernetes.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, archive, namespace, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFromArchive", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromArchive indicates an expected call of DeleteFromArchive.
func (mr *MockChartApplierMockRecorder) DeleteFromArchive(ctx, archive, namespace, name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, archive, namespace, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromArchive", reflect.TypeOf((*MockChartApplier)(nil).DeleteFromArchive), varargs...)
}

// DeleteFromEmbeddedFS mocks base method.
func (m *MockChartApplier) DeleteFromEmbeddedFS(ctx context.Context, embeddedFS embed.FS, chartPath, namespace, name string, opts ...kubernetes.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, embeddedFS, chartPath, namespace, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFromEmbeddedFS", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromEmbeddedFS indicates an expected call of DeleteFromEmbeddedFS.
func (mr *MockChartApplierMockRecorder) DeleteFromEmbeddedFS(ctx, embeddedFS, chartPath, namespace, name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, embeddedFS, chartPath, namespace, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromEmbeddedFS", reflect.TypeOf((*MockChartApplier)(nil).DeleteFromEmbeddedFS), varargs...)
}

// RenderArchive mocks base method.
func (m *MockChartApplier) RenderArchive(archive []byte, releaseName, namespace string, values any) (*chartrenderer.RenderedChart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderArchive", archive, releaseName, namespace, values)
	ret0, _ := ret[0].(*chartrenderer.RenderedChart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderArchive indicates an expected call of RenderArchive.
func (mr *MockChartApplierMockRecorder) RenderArchive(archive, releaseName, namespace, values any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderArchive", reflect.TypeOf((*MockChartApplier)(nil).RenderArchive), archive, releaseName, namespace, values)
}

// RenderEmbeddedFS mocks base method.
func (m *MockChartApplier) RenderEmbeddedFS(embeddedFS embed.FS, chartPath, releaseName, namespace string, values any) (*chartrenderer.RenderedChart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderEmbeddedFS", embeddedFS, chartPath, releaseName, namespace, values)
	ret0, _ := ret[0].(*chartrenderer.RenderedChart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderEmbeddedFS indicates an expected call of RenderEmbeddedFS.
func (mr *MockChartApplierMockRecorder) RenderEmbeddedFS(embeddedFS, chartPath, releaseName, namespace, values any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderEmbeddedFS", reflect.TypeOf((*MockChartApplier)(nil).RenderEmbeddedFS), embeddedFS, chartPath, releaseName, namespace, values)
}
