package entity

type Subject struct {
	Title string `json:"title"`
	BeginTime int `json:"course_bgtime"`
	EndTime int `json:"course_endtime"`
	SoldCount int `json:"sold_count"`
	MinPrice int `json:"course_min_price"`
	MaxPrice int `json:"course_max_price"`
	SumPrice int
}

type SubjectAllData struct {
	Result Subjects `json:"result"`
}
type Subjects struct {
	SubjectList []Subject `json:"sys_course_pkg_list"`
	Lectures Lectures `json:"spe_course_list"`
}


type Lectures struct{
	LectureList []Lecture `json:"data"`
}
type Lecture struct{
	Name string `json:"name"`
	TimePlan string `json:"time_plan"`
	ApplyNum int `json:"apply_num"`
}
