package golibre

import (
	"context"
	"net/http"
	"time"

	"github.com/equalsgibson/golibre/golibre/internal"
)

type ConnectionService struct {
	client *client
	store  *internal.SimpleStore[PatientID, []GraphGlucoseMeasurement]
}

func (c *ConnectionService) GetAllConnectionData(ctx context.Context) ([]ConnectionData, error) {
	endpoint := "/llu/connections"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}

	target := GetAllConnectionsResponse{}
	if err := c.client.Do(req, &target); err != nil {
		return nil, err
	}

	return target.Data, nil
}

func (c *ConnectionService) GetConnectionGraph(ctx context.Context, patientID PatientID) (ConnectionGraphData, error) {
	endpoint := "/llu/connections/" + string(patientID) + "/graph"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		http.NoBody,
	)
	if err != nil {
		return ConnectionGraphData{}, err
	}

	target := ConnectionGraphResponse{}
	if err := c.client.Do(req, &target); err != nil {
		return ConnectionGraphData{}, err
	}

	return target.Data, nil
}

func (c *ConnectionService) RegisterConnectionInStore(ctx context.Context, patientID PatientID, updateInterval time.Duration) error {
	patientData, err := c.GetConnectionGraph(ctx, patientID)
	if err != nil {
		return err
	}

	newConnection := map[PatientID][]GraphGlucoseMeasurement{
		patientData.Connection.PatientID: patientData.GraphData,
	}

	c.store.Set(newConnection)

	return nil
}

func (c *ConnectionService) UnregisterConnectionInStore(ctx context.Context, patientID PatientID) {
	c.store.Evict(patientID)
}

func (c *ConnectionService) pollRegisteredConnections(ctx context.Context) error {
	registeredConnections := c.store.GetAll(ctx)
	updatedData := make(map[PatientID][]GraphGlucoseMeasurement, len(registeredConnections))

	for registeredConnection := range registeredConnections {
		data, err := c.GetConnectionGraph(ctx, registeredConnection)
		if err != nil {
			return err
		}

		updatedData[registeredConnection] = data.GraphData
	}

	c.store.Set(updatedData)

	return nil
}

type ConnectionGraphResponse BaseResponse[ConnectionGraphData]

type GetAllConnectionsResponse BaseResponse[[]ConnectionData]

type ConnectionGraphData struct {
	Connection    ConnectionData            `json:"connection"`
	ActiveSensors []SensorDevicePair        `json:"activeSensors"`
	GraphData     []GraphGlucoseMeasurement `json:"graphData"`
}

type SensorDevicePair struct {
	Sensor Sensor        `json:"sensor"`
	Device PatientDevice `json:"device"`
}

type GraphGlucoseMeasurement struct {
	FactoryTimestamp Timestamp `json:"FactoryTimestamp"` //nolint:tagliatelle
	Timestamp        Timestamp `json:"Timestamp"`        //nolint:tagliatelle
	Type             uint      `json:"type"`
	ValueInMgPerDl   uint      `json:"ValueInMgPerDl"`   //nolint:tagliatelle
	MeasurementColor uint      `json:"MeasurementColor"` //nolint:tagliatelle
	GlucoseUnits     uint      `json:"GlucoseUnits"`     //nolint:tagliatelle
	Value            float32   `json:"Value"`            //nolint:tagliatelle
	IsHigh           bool      `json:"isHigh"`
	IsLow            bool      `json:"isLow"`
}

type ConnectionData struct {
	ID                 UserID             `json:"id"`
	PatientID          PatientID          `json:"patientId"`
	Country            string             `json:"country"`
	Status             uint               `json:"status"`
	FirstName          string             `json:"firstName"`
	LastName           string             `json:"lastName"`
	TargetLow          uint               `json:"targetLow"`
	TargetHigh         uint               `json:"targetHigh"`
	UnitOfMeasurement  UnitOfMeasurement  `json:"uom"`
	Sensor             Sensor             `json:"sensor"`
	AlarmRules         AlarmRules         `json:"alarmRules"`
	GlucoseMeasurement GlucoseMeasurement `json:"glucoseMeasurement"`
	GlucoseItem        GlucoseMeasurement `json:"glucoseItem"`
	GlucoseAlarm       any                `json:"glucoseAlarm"`
	PatientDevice      PatientDevice      `json:"patientDevice"`
	Created            uint               `json:"created"`
}

type PatientDevice struct {
	DID                 string              `json:"did"`
	DTID                uint                `json:"dtid"`
	V                   string              `json:"v"`
	LL                  uint                `json:"ll"`
	HL                  uint                `json:"hl"`
	U                   uint                `json:"u"`
	FixedLowAlarmValues FixedLowAlarmValues `json:"fixedLowAlarmValues"`
	Alarms              bool                `json:"alarms"`
	FixedLowThreshold   uint                `json:"fixedLowThreshold"`
}

type FixedLowAlarmValues struct {
	MGDL  uint    `json:"mgdl"`
	MMOLL float32 `json:"mmoll"`
}

type GlucoseMeasurement struct {
	FactoryTimestamp Timestamp `json:"FactoryTimestamp"` //nolint:tagliatelle
	Timestamp        Timestamp `json:"Timestamp"`        //nolint:tagliatelle
	Type             uint      `json:"type"`
	ValueInMgPerDl   uint      `json:"ValueInMgPerDl"`   //nolint:tagliatelle
	TrendArrow       uint      `json:"TrendArrow"`       //nolint:tagliatelle
	MeasurementColor uint      `json:"MeasurementColor"` //nolint:tagliatelle
	GlucoseUnits     uint      `json:"GlucoseUnits"`     //nolint:tagliatelle
	Value            float32   `json:"Value"`            //nolint:tagliatelle
	IsHigh           bool      `json:"isHigh"`
	IsLow            bool      `json:"isLow"`
}

type Sensor struct {
	DeviceID     string `json:"deviceId"`
	SerialNumber string `json:"sn"`
	Activated    uint   `json:"a"`
	W            uint   `json:"w"`
	PT           uint   `json:"pt"`
	S            bool   `json:"s"`
	LJ           bool   `json:"lj"`
}

type AlarmRules struct {
	C   bool        `json:"c"`
	H   AlarmRuleH  `json:"h"`
	F   AlarmRule   `json:"f"`
	L   AlarmRule   `json:"l"`
	ND  AlarmRuleND `json:"nd"`
	P   uint        `json:"p"`
	R   uint        `json:"r"`
	STD any         `json:"std"`
}

type AlarmRuleH struct {
	TargetHigh     uint    `json:"th"`
	TargetHighMMoL float32 `json:"thmm"`
	D              uint    `json:"d"`
	F              float32 `json:"f"`
}

type AlarmRule struct {
	TargetHigh     uint    `json:"th"`
	TargetHighMMoL float32 `json:"thmm"`
	D              uint    `json:"d"`
	TargetLow      uint    `json:"tl"`
	TargetLowMMoL  float32 `json:"tlmm"`
}

type AlarmRuleND struct {
	I uint `json:"i"`
	R uint `json:"r"`
	L uint `json:"l"`
}
