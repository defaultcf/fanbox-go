// Code generated by ogen, DO NOT EDIT.

package fanbox

import (
	"context"
	"net/url"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// CreatePost invokes createPost operation.
	//
	// Create post.
	//
	// POST /post.create
	CreatePost(ctx context.Context, request OptCreatePostReq, params CreatePostParams) (CreatePostRes, error)
	// DeletePost invokes deletePost operation.
	//
	// Delete post.
	//
	// POST /post.delete
	DeletePost(ctx context.Context, request OptDeletePostReq, params DeletePostParams) (DeletePostRes, error)
	// GetEditablePost invokes getEditablePost operation.
	//
	// Get post.
	//
	// GET /post.getEditable
	GetEditablePost(ctx context.Context, params GetEditablePostParams) (GetEditablePostRes, error)
	// ListManagedPosts invokes listManagedPosts operation.
	//
	// Get posts belongs with you.
	//
	// GET /post.listManaged
	ListManagedPosts(ctx context.Context, params ListManagedPostsParams) (ListManagedPostsRes, error)
	// UpdatePost invokes updatePost operation.
	//
	// Update post.
	//
	// POST /post.update
	UpdatePost(ctx context.Context, request OptUpdatePostReq, params UpdatePostParams) (UpdatePostRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// CreatePost invokes createPost operation.
//
// Create post.
//
// POST /post.create
func (c *Client) CreatePost(ctx context.Context, request OptCreatePostReq, params CreatePostParams) (CreatePostRes, error) {
	res, err := c.sendCreatePost(ctx, request, params)
	return res, err
}

func (c *Client) sendCreatePost(ctx context.Context, request OptCreatePostReq, params CreatePostParams) (res CreatePostRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/post.create"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreatePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Origin",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Origin))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "User-Agent",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.UserAgent))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securityCsrfToken(ctx, "CreatePost", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"CsrfToken\"")
			}
		}
		{

			switch err := c.securitySessionId(ctx, "CreatePost", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 1
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"SessionId\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000011},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreatePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeletePost invokes deletePost operation.
//
// Delete post.
//
// POST /post.delete
func (c *Client) DeletePost(ctx context.Context, request OptDeletePostReq, params DeletePostParams) (DeletePostRes, error) {
	res, err := c.sendDeletePost(ctx, request, params)
	return res, err
}

func (c *Client) sendDeletePost(ctx context.Context, request OptDeletePostReq, params DeletePostParams) (res DeletePostRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/post.delete"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeDeletePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Origin",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Origin))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "User-Agent",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.UserAgent))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securityCsrfToken(ctx, "DeletePost", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"CsrfToken\"")
			}
		}
		{

			switch err := c.securitySessionId(ctx, "DeletePost", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 1
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"SessionId\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000011},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDeletePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetEditablePost invokes getEditablePost operation.
//
// Get post.
//
// GET /post.getEditable
func (c *Client) GetEditablePost(ctx context.Context, params GetEditablePostParams) (GetEditablePostRes, error) {
	res, err := c.sendGetEditablePost(ctx, params)
	return res, err
}

func (c *Client) sendGetEditablePost(ctx context.Context, params GetEditablePostParams) (res GetEditablePostRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/post.getEditable"
	uri.AddPathParts(u, pathParts[:]...)

	q := uri.NewQueryEncoder()
	{
		// Encode "postId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "postId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.PostId))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Origin",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Origin))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "User-Agent",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.UserAgent))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securitySessionId(ctx, "GetEditablePost", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"SessionId\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeGetEditablePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListManagedPosts invokes listManagedPosts operation.
//
// Get posts belongs with you.
//
// GET /post.listManaged
func (c *Client) ListManagedPosts(ctx context.Context, params ListManagedPostsParams) (ListManagedPostsRes, error) {
	res, err := c.sendListManagedPosts(ctx, params)
	return res, err
}

func (c *Client) sendListManagedPosts(ctx context.Context, params ListManagedPostsParams) (res ListManagedPostsRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/post.listManaged"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Origin",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Origin))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "User-Agent",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.UserAgent))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securitySessionId(ctx, "ListManagedPosts", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"SessionId\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeListManagedPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdatePost invokes updatePost operation.
//
// Update post.
//
// POST /post.update
func (c *Client) UpdatePost(ctx context.Context, request OptUpdatePostReq, params UpdatePostParams) (UpdatePostRes, error) {
	res, err := c.sendUpdatePost(ctx, request, params)
	return res, err
}

func (c *Client) sendUpdatePost(ctx context.Context, request OptUpdatePostReq, params UpdatePostParams) (res UpdatePostRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/post.update"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdatePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Origin",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Origin))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "User-Agent",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.UserAgent))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securitySessionId(ctx, "UpdatePost", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"SessionId\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeUpdatePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
