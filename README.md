# FANBOX GO

```golang
type customSecuritySource struct{}

func (s customSecuritySource) CsrfToken(ctx context.Context, operationName string) (CsrfToken, error) {
	return CsrfToken{
		APIKey: "CsrfToken",
	}, nil
}

func (s customSecuritySource) SessionId(ctx context.Context, operationName string) (SessionId, error) {
	return SessionId{
		APIKey: "SessionId",
	}, nil
}

func main() {
	s := &customSecuritySource{}
	c, err := NewClient("https://api.fanbox.cc", s)
	if err != nil {
		os.Exit(1)
	}

	res, err := c.ListManagedPosts(context.TODO(), ListManagedPostsParams{Origin: "https://www.fanbox.cc"})
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Printf("res: %+v", res)
}
```
