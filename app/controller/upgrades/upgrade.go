package upgrades

import (
	"fmt"
	"mus_projekt/app/model/query"
	"mus_projekt/app/service"
)

func CreateApiData() map[string]interface{} {

	m := make(map[string]interface{})

	fmt.Println("All The Upgrades")
	ush, err := service.GetAllUSH()
	if err != nil {
		m[query.SUCCESS] = false
		return m
	}
	m["upgrades"] = []map[string]interface{}{}
	for _, k := range *ush {
		fmt.Println(k)
		q := make(map[string]interface{})
		q[query.USER_EMAIL_ADDRESS] = k.EmailAddress()
		q[query.SERVICE_UPGRADE_HELPER_ID] = k.GetServiceID()
		q[query.USER_PRIVILEGES] = k.UserAccess()

		m["upgrades"] = append(m["upgrades"].([]map[string]interface{}), q)

	}
	return m

}
