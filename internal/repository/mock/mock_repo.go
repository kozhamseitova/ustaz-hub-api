// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(ctx context.Context, u *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), ctx, u)
}

// DeleteUser mocks base method.
func (m *MockUser) DeleteUser(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUser)(nil).DeleteUser), ctx, id)
}

// GetUserByUsername mocks base method.
func (m *MockUser) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserMockRecorder) GetUserByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUser)(nil).GetUserByUsername), ctx, username)
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(ctx context.Context, u *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), ctx, u)
}

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProduct) CreateProduct(ctx context.Context, p *entity.Product) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, p)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductMockRecorder) CreateProduct(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProduct)(nil).CreateProduct), ctx, p)
}

// DeleteProduct mocks base method.
func (m *MockProduct) DeleteProduct(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductMockRecorder) DeleteProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProduct)(nil).DeleteProduct), ctx, id)
}

// GetAllProducts mocks base method.
func (m *MockProduct) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts", ctx)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductMockRecorder) GetAllProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProduct)(nil).GetAllProducts), ctx)
}

// GetProductById mocks base method.
func (m *MockProduct) GetProductById(ctx context.Context, id int64) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", ctx, id)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductMockRecorder) GetProductById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProduct)(nil).GetProductById), ctx, id)
}

// GetProductsByUserId mocks base method.
func (m *MockProduct) GetProductsByUserId(ctx context.Context, userId int64) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByUserId", ctx, userId)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByUserId indicates an expected call of GetProductsByUserId.
func (mr *MockProductMockRecorder) GetProductsByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByUserId", reflect.TypeOf((*MockProduct)(nil).GetProductsByUserId), ctx, userId)
}

// UpdateProduct mocks base method.
func (m *MockProduct) UpdateProduct(ctx context.Context, p *entity.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductMockRecorder) UpdateProduct(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProduct)(nil).UpdateProduct), ctx, p)
}

// MockPost is a mock of Post interface.
type MockPost struct {
	ctrl     *gomock.Controller
	recorder *MockPostMockRecorder
}

// MockPostMockRecorder is the mock recorder for MockPost.
type MockPostMockRecorder struct {
	mock *MockPost
}

// NewMockPost creates a new mock instance.
func NewMockPost(ctrl *gomock.Controller) *MockPost {
	mock := &MockPost{ctrl: ctrl}
	mock.recorder = &MockPostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPost) EXPECT() *MockPostMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPost) CreatePost(ctx context.Context, p *entity.Post) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, p)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostMockRecorder) CreatePost(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPost)(nil).CreatePost), ctx, p)
}

// DeletePost mocks base method.
func (m *MockPost) DeletePost(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostMockRecorder) DeletePost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPost)(nil).DeletePost), ctx, id)
}

// GetAllPosts mocks base method.
func (m *MockPost) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPosts", ctx)
	ret0, _ := ret[0].([]entity.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPosts indicates an expected call of GetAllPosts.
func (mr *MockPostMockRecorder) GetAllPosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPosts", reflect.TypeOf((*MockPost)(nil).GetAllPosts), ctx)
}

// GetPostById mocks base method.
func (m *MockPost) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostById", ctx, id)
	ret0, _ := ret[0].(*entity.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostById indicates an expected call of GetPostById.
func (mr *MockPostMockRecorder) GetPostById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostById", reflect.TypeOf((*MockPost)(nil).GetPostById), ctx, id)
}

// GetPostsByUserId mocks base method.
func (m *MockPost) GetPostsByUserId(ctx context.Context, userId int64) ([]entity.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByUserId", ctx, userId)
	ret0, _ := ret[0].([]entity.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByUserId indicates an expected call of GetPostsByUserId.
func (mr *MockPostMockRecorder) GetPostsByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByUserId", reflect.TypeOf((*MockPost)(nil).GetPostsByUserId), ctx, userId)
}

// UpdatePost mocks base method.
func (m *MockPost) UpdatePost(ctx context.Context, p *entity.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostMockRecorder) UpdatePost(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPost)(nil).UpdatePost), ctx, p)
}

// MockComment is a mock of Comment interface.
type MockComment struct {
	ctrl     *gomock.Controller
	recorder *MockCommentMockRecorder
}

// MockCommentMockRecorder is the mock recorder for MockComment.
type MockCommentMockRecorder struct {
	mock *MockComment
}

// NewMockComment creates a new mock instance.
func NewMockComment(ctrl *gomock.Controller) *MockComment {
	mock := &MockComment{ctrl: ctrl}
	mock.recorder = &MockCommentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComment) EXPECT() *MockCommentMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockComment) CreateComment(ctx context.Context, p *entity.CreateComment) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, p)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentMockRecorder) CreateComment(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockComment)(nil).CreateComment), ctx, p)
}

// GetAllPostsComments mocks base method.
func (m *MockComment) GetAllPostsComments(ctx context.Context, postId int64) ([]entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPostsComments", ctx, postId)
	ret0, _ := ret[0].([]entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPostsComments indicates an expected call of GetAllPostsComments.
func (mr *MockCommentMockRecorder) GetAllPostsComments(ctx, postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPostsComments", reflect.TypeOf((*MockComment)(nil).GetAllPostsComments), ctx, postId)
}

// GetAllProductsComments mocks base method.
func (m *MockComment) GetAllProductsComments(ctx context.Context, productId int64) ([]entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProductsComments", ctx, productId)
	ret0, _ := ret[0].([]entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProductsComments indicates an expected call of GetAllProductsComments.
func (mr *MockCommentMockRecorder) GetAllProductsComments(ctx, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProductsComments", reflect.TypeOf((*MockComment)(nil).GetAllProductsComments), ctx, productId)
}

// GetCommentsByParentId mocks base method.
func (m *MockComment) GetCommentsByParentId(ctx context.Context, parentId int64) ([]entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsByParentId", ctx, parentId)
	ret0, _ := ret[0].([]entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByParentId indicates an expected call of GetCommentsByParentId.
func (mr *MockCommentMockRecorder) GetCommentsByParentId(ctx, parentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByParentId", reflect.TypeOf((*MockComment)(nil).GetCommentsByParentId), ctx, parentId)
}