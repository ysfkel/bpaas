package organization_request

type IOrganizationRequest interface {
	Create(dataPath string) error
}
