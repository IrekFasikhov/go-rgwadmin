# go-rgwadmin

Package radosgw-admin wraps http://docs.ceph.com/docs/mimic/radosgw/adminops

Requires Go 1.11 or newer. Tested with Liminous and Mimic releases!


[Example App:](https://github.com/Kataklysm/go-rgwadmin/examples/)

  ```go
  func main() {

	endpoint := os.Getenv("RGW_ENDPOINT")
	accessKey := os.Getenv("RGW_ACCESS_KEY")
	secretKey := os.Getenv("RGW_SECRET_KEY")
    
  .........

	api, err := rgwadmin.NewRGW(endpoint, accessKey, secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *getUser {
		u := rgwadmin.User{
			UID: *uid,
		}
		user, _ := api.GetUser(u)
		fmt.Println(user)
	}
	if *createUser {
		u := rgwadmin.User{
			UID:         *uid,
			DisplayName: *displayName,
			Email:       *email,
			Suspended:   suspended,
		}
		fmt.Println(api.CreateUser(u))
	}
  .....
  ```
