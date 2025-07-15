package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ajustc/awb-mrt-schedules/common/client"
)

type Service interface {
	GetAll() (response []StationResponse, err error)
	GetByID(id string) (response []ScheduleResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *service) GetAll() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return response, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return response, err
	}

	for _, item := range stations {
		response = append(response, StationResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return
}

func (s *service) GetByID(id string) (response []ScheduleResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return response, err
	}

	var schedule []Schedule
	err = json.Unmarshal(byteResponse, &schedule)
	if err != nil {
		return response, err
	}

	var scheduleSelected Schedule
	for _, item := range schedule {
		if item.StationID == id {
			scheduleSelected = item
			break
		}
	}

	if scheduleSelected.StationID == "" {
		err = errors.New("schedule not found")
		return
	}

	response, err = ConvertDataToResponse(scheduleSelected)
	if err != nil {
		return response, err
	}

	return
}

func ConvertDataToResponse(schedule Schedule) (response []ScheduleResponse, err error) {
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus"
		BundaranHITripName = "Stasiun Bundaran HI"
	)

	scheduleLebakBulus := schedule.ScheduleLB
	scheduleBundaranHI := schedule.ScheduleHI

	scheduleLebakBulusParsed, err := ConvertSchduleToTimeFormat(scheduleLebakBulus)
	if err != nil {
		return response, err
	}

	scheduleBundaranHIParsed, err := ConvertSchduleToTimeFormat(scheduleBundaranHI)
	if err != nil {
		return response, err
	}

	for _, item := range scheduleLebakBulusParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: LebakBulusTripName,
				Time:        item.Format("15:04"),
			})
		}
	}

	for _, item := range scheduleBundaranHIParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: BundaranHITripName,
				Time:        item.Format("15:04"),
			})
		}
	}

	return
}

func ConvertSchduleToTimeFormat(schduled string) (response []time.Time, err error) {
	var (
		parsedTime time.Time
		schedules  = strings.Split(schduled, ",")
	)
	for _, item := range schedules {
		trimTime := strings.TrimSpace(item)
		if trimTime == "" {
			continue
		}

		parsedTime, err = time.Parse("15:04", trimTime)
		if err != nil {
			err = errors.New("invalid time format : " + trimTime)
			return response, err
		}

		response = append(response, parsedTime)
	}
	return
}
