package calendar

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	ProdId = "-//github.com//nxm"
)

var (
	CalendarEvents []VEvent
)

type VCalendar struct {
	ProdId   string  `str:"PRODID"`
	Version  float32 `str:"VERSION"`
	CalScale bool    `str:"CALSCALE"`
}

type VEvent struct {
	DTStamp     string    `str:"DTSTAMP"`
	DTStart     string    `str:"DTSTART"`
	DTEnd       string    `str:"DTEND"`
	Summary     string    `str:"SUMMARY"`
	UID         uuid.UUID `str:"UID"`
	Location    string    `str:"LOCATION"`
	Description string    `str:"DESCRIPTION"`
}

func (VCalendar) AddEvent(event VEvent) {
	CalendarEvents = append(CalendarEvents, event)
}

func (VCalendar) GetEvents() []VEvent {
	return CalendarEvents
}

func CreateCalendar() VCalendar {
	return VCalendar{ProdId: ProdId, Version: 2.0, CalScale: false}
}

func CreateEvent(start, end time.Time) VEvent {
	event := VEvent{
		DTStamp: ParseTimeToCalendar(time.Now()),
		DTStart: ParseTimeToCalendar(start),
		DTEnd:   ParseTimeToCalendar(end),
		UID:     GenerateUID(),
	}

	return event
}

func (e *VEvent) SetTitle(title string) {
	e.Summary = title
}

func (e *VEvent) SetDescription(description string) {
	e.Description = description
}

func (e *VEvent) SetLocation(location string) {
	e.Location = location
}

func ParseTimeToCalendar(time time.Time) string {
	str := fmt.Sprintf("%d%02d%02dT%02d%02d%02d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
	return str
}

func GenerateUID() uuid.UUID {
	u, _ := uuid.NewRandom()
	return u
}

func (cal VCalendar) Save(fileName string) error {
	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	for _, outLines := range cal.ParseVCalendar() {
		_, err := f.WriteString(outLines)

		if err != nil {
			return err
		}
	}

	f.Sync()

	return nil
}

func (cal VCalendar) ParseVCalendar() []string {
	var strs []string

	s := reflect.ValueOf(&cal).Elem()
	typeOfT := s.Type()

	strs = append(strs, "BEGIN:VCALENDAR\n")

	for i := 0; i < s.NumField(); i++ {
		var value string
		f := s.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())

		//shitty fix, sorry for that
		if typeOfT.Field(i).Type.String() == "float32" {
			value = fmt.Sprintf("%f", f.Interface().(float32))
		} else if typeOfT.Field(i).Type.String() == "bool" {
			value = strconv.FormatBool(f.Interface().(bool))
		} else {
			value = f.Interface().(string)
		}

		strs = append(strs, typeOfT.Field(i).Name+":"+value+"\n")

	}

	for _, evn := range cal.GetEvents() {
		tempEvent := cal.ParseVEvent(evn)
		strs = append(strs, tempEvent...)
	}

	strs = append(strs, "END:VCALENDAR")

	return strs
}

func (VCalendar) ParseVEvent(ev VEvent) []string {
	var strs []string

	//i cant cast parameter from interface{} to VEvent cuz
	//reflect: call of reflect.Value.NumField on interface Value
	s := reflect.ValueOf(&ev).Elem()
	typeOfT := s.Type()

	//beginning event
	strs = append(strs, "BEGIN:VEVENT\n")

	for i := 0; i < s.NumField(); i++ {
		var value string
		f := s.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())

		//shitty fix, rly sorry
		if typeOfT.Field(i).Type.String() == "uuid.UUID" {
			value = f.Interface().(uuid.UUID).String()
		} else {
			value = f.Interface().(string)
		}

		//panic: interface conversion: interface {} is uuid.UUID, not string
		//strs = append(strs, typeOfT.Field(i).Name + ":" + f.Interface().(string) + "\n")

		strs = append(strs, typeOfT.Field(i).Name+":"+value+"\n")
	}

	//ending event
	strs = append(strs, "END:VEVENT\n")

	return strs
}
