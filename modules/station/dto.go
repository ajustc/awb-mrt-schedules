package station

type Station struct {
	ID   string `json:"nid"`
	Name string `json:"title"`
}

type StationResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	StationID   string `json:"nid"`
	StationName string `json:"title"`
	ScheduleHI  string `json:"jadwal_hi_biasa"`
	ScheduleLB  string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	StationName string `json:"station"`
	Time        string `json:"time"`
}
