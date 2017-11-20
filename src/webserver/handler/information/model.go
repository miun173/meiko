package information

const (
	TableNameInformation = "informations"
)

type informationResponse struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type getSummaryResponse struct {
	Last   []informationResponse `json:"last"`
	Recent []informationResponse `json:"recent"`
}
type detailInfromationParams struct {
	ID string
}
type detailInfromationArgs struct {
	ID int64
}

type createParams struct {
	Title       string
	Description string
	ScheduleID  string
	FilesID     string
}
type createArgs struct {
	Title       string
	Description string
	ScheduleID  int64
	FilesID     []string
}
type updateParams struct {
	ID          string
	Title       string
	Description string
	ScheduleID  string
}

type upadateArgs struct {
	ID          int64
	Title       string
	Description string
	ScheduleID  int64
}

type deleteParams struct {
	ID string
}

type deleteArgs struct {
	ID int64
}

type readListParams struct {
	Total string
	Page  string
}
type readListArgs struct {
	Total uint16
	Page  uint16
}
