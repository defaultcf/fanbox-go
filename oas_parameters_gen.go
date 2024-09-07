// Code generated by ogen, DO NOT EDIT.

package fanbox

// CreatePostParams is parameters of createPost operation.
type CreatePostParams struct {
	Origin string
	// You must set user agent.
	UserAgent string
}

// DeletePostParams is parameters of deletePost operation.
type DeletePostParams struct {
	Origin string
	// You must set user agent.
	UserAgent string
}

// GetEditablePostParams is parameters of getEditablePost operation.
type GetEditablePostParams struct {
	Origin string
	// You must set user agent.
	UserAgent string
	// Post's id.
	PostId string
}

// ListManagedPostsParams is parameters of listManagedPosts operation.
type ListManagedPostsParams struct {
	Origin string
	// You must set user agent.
	UserAgent string
}

// UpdatePostParams is parameters of updatePost operation.
type UpdatePostParams struct {
	Origin string
	// You must set user agent.
	UserAgent string
}
