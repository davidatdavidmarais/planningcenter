# planningcenter
Planning Center API Integration



## OAuth
### Exchanging oauth redirection code for access token

```
res, err := pc.OAuthClient.ExchangeToken("auth_code", "redirect_url")
if err != nil {
	return err
}
```

### Refreshing oauth access token

```
res, err := pc.OAuthClient.RefreshToken("refresh_token")
if err != nil {
	return err
}
```


