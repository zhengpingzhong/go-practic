// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Client   string `json:"client"`
}

type LoginRes struct {
	Msg              string `json:"msg"`
	Code             string `json:"code"`
	OrganizationName string `json:"organizationName"`
	Roleid           string `json:"roleid"`
	Mobile           string `json:"mobile"`
	UserName         string `json:"userName"`
	UserId           string `json:"userId"`
	GrandpaId        string `json:"grandpaId"`
	Token            string `json:"token"`
	OrganizationId   string `json:"organizationId"`
	ParentName       string `json:"parentName"`
	ParentOrgId      string `json:"parentOrgId"`
	OrgCode          string `json:"orgCode"`
	Expire           string `json:"expire"`
	StaffName        string `json:"staffName"`
	GrandpaName      string `json:"grandpaName"`
	StaffId          string `json:"staffId"`
}
