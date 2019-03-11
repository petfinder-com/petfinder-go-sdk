package pfapi

type OrganizationResponse struct {
	Organizations []Organization
	Pagination    Pagination
}

type Organization struct {
	ID               string
	Name             string
	Email            string
	Phone            string
	Address          Address
	Hours            Hours
	URL              string
	WebSite          string
	MissionStatement string
	AdoptionPolicy   AdoptionPolicy
	SocialMedia      SocialMedia `mapstructure:"social_media"`
	Photos           []Photo
	Links            OrganizationLinks `mapstructure:"_links"`
}

type Hours struct {
	Monday    string
	Tuesday   string
	Wednesday string
	Thursday  string
	Friday    string
	Saturday  string
	Sunday    string
}

type SocialMedia struct {
	Facebook  string
	Twitter   string
	Youtube   string
	Instagram string
	Pinterest string
}

type OrganizationLinks struct {
	Self    Link
	Animals Link
}

type AdoptionPolicy struct {
	Policy string
	URL    string
}
