package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/jessevdk/go-flags"

	"github.com/1024pix/go-metabase/metabase"
	"github.com/1024pix/go-metabase/metabaseutil"
	mo "github.com/grokify/goauth/metabase"
)

type Options struct {
	Env string `short:"e" long:"envfile" description:"Env File" required:"false"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	read, err := config.LoadDotEnv([]string{
		opts.Env, os.Getenv("ENV_PATH"), ".env"}, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(read)

	clientCfg := mo.Config{
		BaseURL:       os.Getenv("METABASE_BASE_URL"),
		Username:      os.Getenv("METABASE_USERNAME"),
		Password:      os.Getenv("METABASE_PASSWORD"),
		SessionID:     os.Getenv("METABASE_SESSION_ID"),
		TLSSkipVerify: stringsutil.ToBool(os.Getenv("METABASE_TLS_SKIP_VERIFY"))}

	err = clientCfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	apiClient, _, err := metabaseutil.NewApiClient(clientCfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("list groups :")

	err = listPermissionsGroup(apiClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("get group :")
	err = getTestGroup(apiClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}

func listPermissionsGroup(apiClient *metabase.APIClient) error {

	request := apiClient.PermissionsApi.ListPermissionsGroups(context.Background())

	permissionsGroup, resp, err := apiClient.PermissionsApi.ListPermissionsGroupsExecute(request)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("Status Code [%v]", resp.StatusCode)
	}

	for _, permission := range permissionsGroup {
		fmt.Printf("PERMISSION_ID [%v] PERMISSION_NAME [%v] PERMISSION_MEMBER_COUNT [%v]\n", *permission.Id, *permission.Name, *permission.MemberCount)

		if permission.Members == nil {
			fmt.Println("    ->No members")
			continue
		}

		for _, user := range *permission.Members {
			fmt.Printf("    - User : %v\n", *user.Email)
		}
	}
	return nil
}

func getTestGroup(apiClient *metabase.APIClient) error {

	request := apiClient.PermissionsApi.GetPermissionsGroup(context.Background(), 5)

	permissionsGroup, resp, err := apiClient.PermissionsApi.GetPermissionsGroupExecute(request)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("Status Code [%v]", resp.StatusCode)
	}

	// if permissionsGroup == nil {
	//     return fmt.Errorf("GetPermissionsGroup returned nil object")
	// }

	id := int32(0)
	if permissionsGroup.Id != nil {
		id = *permissionsGroup.Id
	}
	name := ""
	if permissionsGroup.Name != nil {
		name = *permissionsGroup.Name
	}
	memberCount := int32(0)
	if permissionsGroup.MemberCount != nil {
		memberCount = *permissionsGroup.MemberCount
	}

	fmt.Printf("PERMISSION_ID [%v] PERMISSION_NAME [%v] PERMISSION_MEMBER_COUNT [%v]\n", id, name, memberCount)

	if permissionsGroup.Members == nil {
		fmt.Println("    - No members")
		return nil
	}

	for _, user := range *permissionsGroup.Members {
		fmt.Printf("    - User : %v\n", *user.Email)
	}
	return nil
}
