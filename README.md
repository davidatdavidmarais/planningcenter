# planningcenter
Planning Center API Integration

### Initialization

```
// Where os.Getenv("PLANNING_CENTER_CLIENT_ID") & os.Getenv("PLANNING_CENTER_CLIENT_SECRET") should be retrieved from your planning center developer dashboard.
pc := planningcenter.New(
	os.Getenv("PLANNING_CENTER_CLIENT_ID"),
	os.Getenv("PLANNING_CENTER_CLIENT_SECRET"),
)
```


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
## People
## Getting the user details for an access token

```
res, err := pc.PeopleClient.Me("access_token")
if err != nil {
	return err
}
```

## Services
### Getting schedules

```
// person_id retrieved from pc.PeopleClient.Me("access_token")
res, err := pc.ServicesClient.Schdules("access_token", "person_id")
if err != nil {
	return err
}
```


## Getting items for a schedule

```
// service_type_id & plan_id retrieved from pc.ServicesClient.Schdules("access_token", "person_id")
res, err := pc.ServicesClient.Items("access_token", "service_type_id", "plan_id")
if err != nil {
	return err
}
```

## Getting an arrangement for a song item

```
// service_type_id & plan_id retrieved from pc.ServicesClient.Schdules("access_token", "person_id")
// item_id retrieved from pc.ServicesClient.Items("access_token", "service_type_id", "plan_id")
res, err := pc.ServicesClient.Arrangement("access_token", "service_type_id", "plan_id", "item_id")
if err != nil {
	return err
}
```
