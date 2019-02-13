package main

import (
	"flag"
	"fmt"
	"os"

	rgwadmin "github.com/Kataklysm/golang-radosgw/"
)

func main() {

	endpoint := os.Getenv("RGW_ENDPOINT")
	accessKey := os.Getenv("RGW_ACCESS_KEY")
	secretKey := os.Getenv("RGW_SECRET_KEY")

	getUser := flag.Bool("GetUser", false, "")
	createUser := flag.Bool("GreateUser", false, "")
	removeUser := flag.Bool("RemoveUser", false, "")
	modifyUser := flag.Bool("ModifyUser", false, "")
	quota := flag.Bool("Quota", false, "")
	getPolicy := flag.Bool("GetPolicy", false, "")
	getUsage := flag.Bool("GetUsage", false, "")
	trimUsage := flag.Bool("TrimUsage", false, "")

	ListBuckets := flag.Bool("listBuckets", false, "")
	InfoBucket := flag.Bool("infoBucket", false, "")
	bucketName := flag.String("bucketName", "", "")

	uid := flag.String("uid", "", "username")
	displayName := flag.String("displayName", "", "Display Name")
	email := flag.String("email", "", "Email")
	suspended := flag.Int("suspend", 0, "suspend")
	maxSize := flag.Uint64("max-size", 0, "max-size KB")

	flag.Parse()

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

	if *removeUser {
		u := rgwadmin.User{
			UID: *uid,
		}
		api.RemoveUser(u)
	}

	if *modifyUser {
		u := rgwadmin.User{
			UID:         *uid,
			Suspended:   suspended,
			MaxBuckets:  toIntPtr(10),
			GenerateKey: toBoolPtr(false),
			Keys: []rgwadmin.Key{
				{
					AccessKey: "sdfsdfsdfsdfsdfsdfsdfsdSDFSDF",
					SecretKey: "sfsdfsdfsfsdfsfwer2r2wfr423refd",
				},
			},
		}
		fmt.Println(api.ModifyUser(u))
	}

	if *quota {
		q := rgwadmin.Quota{
			UID:        *uid,
			Enabled:    toBoolPtr(true),
			MaxObjects: toUInt64Ptr(100000000),
			MaxSizeKb:  maxSize,
		}

		q2, _ := api.SetQuota(q, "bucket")
		fmt.Println(q2)
	}

	if *ListBuckets {
		fmt.Println(api.ListBuckets())
	}

	if *InfoBucket {
		fmt.Println(api.InfoBucket(*bucketName))
	}

	if *getPolicy {
		fmt.Println(api.GetBucketPolicy(*bucketName))
	}

	if *getUsage {
		usage, _ := api.GetUsage(rgwadmin.Usage{
			// ShowSummary: toBoolPtr(false),
		})
		fmt.Println(usage)
	}

	if *trimUsage {
		api.TrimUsage(
			rgwadmin.Usage{
				RemoveAll: toBoolPtr(true),
			},
		)
	}
}

func toIntPtr(x int) *int {
	return &x
}

func toUInt64Ptr(x uint64) *uint64 {
	return &x
}
func toBoolPtr(x bool) *bool {
	return &x
}
