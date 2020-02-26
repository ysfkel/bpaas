package organization

type IOrganization interface {
	List(companyName string) error
}
