# planningcenter
Go wrapper for Planning Center API

### Initialization

> Retrieve the **client_id** and **client_secret** from your *planning center developer dashboard*
```
pc := planningcenter.New(
	os.Getenv("PLANNING_CENTER_CLIENT_ID"),
	os.Getenv("PLANNING_CENTER_CLIENT_SECRET"),
)
```

### OAuth
#### Exchanging oauth redirection code for access token

```
res, err := pc.OAuthClient.ExchangeToken("auth_code", "redirect_url")
if err != nil {
	return err
}
```

#### Refreshing oauth access token

```
res, err := pc.OAuthClient.RefreshToken("refresh_token")
if err != nil {
	return err
}
```
### People
#### Getting the user details for an access token

```
res, err := pc.PeopleClient.Me("access_token")
if err != nil {
	return err
}
```

### Services
#### Getting schedules
> Retrieve **person_id** from *pc.PeopleClient.Me("access_token")*
```
res, err := pc.ServicesClient.Schdules("access_token", "person_id")
if err != nil {
	return err
}
```


#### Getting items for a schedule

> Retrieve **service_type_id** & **plan_id** from *pc.ServicesClient.Schdules("access_token", "person_id")*
```
res, err := pc.ServicesClient.Items("access_token", "service_type_id", "plan_id")
if err != nil {
	return err
}
```

#### Getting an arrangement for a song item
> Retrieve **service_type_id** & **plan_id** from *pc.ServicesClient.Schdules("access_token", "person_id")*
> 
> Retrieve **item_id** from *pc.ServicesClient.Items("access_token", "service_type_id", "plan_id")*
```
res, err := pc.ServicesClient.Arrangement("access_token", "service_type_id", "plan_id", "item_id")
if err != nil {
	return err
}
```
