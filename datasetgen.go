package datasetgen

import (
	"encoding/json"
	"regexp"
	"time"
)

type (
	ContactData struct {
		Day       *string `json:"日付"`
		DayOfWeek *int    `json:"曜日"`
		Nine13    *int    `json:"9-13時"`
		One317    *int    `json:"13-17時"`
		One721    *int    `json:"17-21時"`
		Date      *string `json:"date"`
		W         *int    `json:"w"`
		ShortDate *string `json:"short_date"`
		SubTotal  *int    `json:"小計"`
	}

	QuerentData struct {
		Day       *string `json:"日付"`
		DayOfWeek *int    `json:"曜日"`
		Nine17    *int    `json:"9-17時"`
		One79     *int    `json:"17-翌9時"`
		Date      *string `json:"date"`
		W         *int    `json:"w"`
		ShortDate *string `json:"short_date"`
		SubTotal  *int    `json:"小計"`
	}

	PatientData struct {
		ReleaseDate *string `json:"リリース日"`
		Address     *string `json:"居住地"`
		Age         *string `json:"年代"`
		Sex         *string `json:"性別"`
		Leave       *string `json:"退院"`
		Date        *string `json:"date"`
	}

	PatientsSummaryData struct {
		Day      *string `json:"日付"`
		SubTotal *int    `json:"小計"`
	}

	DischargesSummaryData struct {
		Day      *string `json:"日付"`
		SubTotal *int    `json:"小計"`
	}

	InspectionsSummaryData struct {
		InAomori []*int `json:"県内"`
		Others   []*int `json:"その他"`
	}

	InspectionPersonsDataset struct {
		Label *string `json:"label"`
		Data  []*int  `json:"data"`
	}

	InspectionStatusSummaryChildrenChildren struct {
		Attr  *string `json:"attr"`
		Value *int    `json:"value"`
	}

	InspectionStatusSummaryChildren struct {
		Attr     *string                                   `json:"attr"`
		Value    *int                                      `json:"value"`
		Children []InspectionStatusSummaryChildrenChildren `json:"children"`
	}

	Contacts struct {
		Date *string       `json:"date"`
		Data []ContactData `json:"data"`
	}

	Querents struct {
		Date *string       `json:"date"`
		Data []QuerentData `json:"data"`
	}

	Patients struct {
		Date *string       `json:"date"`
		Data []PatientData `json:"data"`
	}

	PatientsSummary struct {
		Date *string               `json:"date"`
		Data []PatientsSummaryData `json:"data"`
	}

	DischargesSummary struct {
		Date *string                 `json:"date"`
		Data []DischargesSummaryData `json:"data"`
	}

	InspectionsSummary struct {
		Date   *string                `json:"date"`
		Data   InspectionsSummaryData `json:"data"`
		Labels []*string              `json:"labels"`
	}

	InspectionPersons struct {
		Date     *string                    `json:"date"`
		Labels   []*string                  `json:"labels"`
		Datasets []InspectionPersonsDataset `json:"datasets"`
	}

	InspectionStatusSummary struct {
		Date     *string                           `json:"date"`
		Attr     *string                           `json:"attr"`
		Value    *int                              `json:"value"`
		Children []InspectionStatusSummaryChildren `json:"children"`
	}

	DataSet struct {
		Contacts                Contacts                `json:"contacts"`
		Querents                Querents                `json:"querents"`
		Patients                Patients                `json:"patients"`
		PatientsSummary         PatientsSummary         `json:"patients_summary"`
		DischargesSummary       DischargesSummary       `json:"discharges_summary"`
		InspectionsSummary      InspectionsSummary      `json:"inspections_summary"`
		InspectionPersons       InspectionPersons       `json:"inspection_persons"`
		InspectionStatusSummary InspectionStatusSummary `json:"inspection_status_summary"`
		LastUpdate              *string                 `json:"lastUpdate"`
	}
)

func NewDataSet() *DataSet {
	return new(DataSet)
}

func (d *DataSet) SetLastUpdateDate() {
	t := time.Now().In(time.Local).Format("2006/01/02 15:04")
	d.LastUpdate = &t
}

func (d *DataSet) GenerateJson() (string, error) {
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (d *DataSet) SetPatients(webapi *WebApi) error {
	for _, resource := range webapi.Result.Resources {
		if regexp.MustCompile(`^\d{8}_陽性患者関係.csv$`).MatchString(resource.Name) {

			// Get csv
			patientsStatus := NewPatientsStatus()
			if err := patientsStatus.Get(resource.URL); err != nil {
				return err
			}

			// Set Dataset
			for _, v := range *patientsStatus {
				releaseDate, err := TransformRFC3339Z(v.PublishDate)
				if err != nil {
					return err
				}

				date, err := Transform20060102(v.FixedDate)
				if err != nil {
					return err
				}
				address := v.Address
				age := v.Age
				sex := v.Sex

				d.Patients.Data = append(
					d.Patients.Data,
					PatientData{
						ReleaseDate: &releaseDate,
						Address:     &address,
						Age:         &age,
						Sex:         &sex,
						Date:        &date,
					},
				)
			}

			date := Transform200601021504(resource.Updated)
			d.Patients.Date = &date
		}
	}

	return nil
}

func (d *DataSet) SetPatientsSummary(webapi *WebApi) error {
	for _, resource := range webapi.Result.Resources {
		if regexp.MustCompile(`^\d{8}_検査実施状況.csv$`).MatchString(resource.Name) {
			// Get csv
			inspectionsStatus := NewInspectionsStatus()
			if err := inspectionsStatus.Get(resource.URL); err != nil {
				return err
			}

			// Set Dataset
			for _, v := range *inspectionsStatus {
				dateRFC3339Z, err := TransformRFC3339Z(v.Date)
				if err != nil {
					return err
				}

				positive := v.Positive

				d.PatientsSummary.Data = append(
					d.PatientsSummary.Data,
					PatientsSummaryData{
						Day:      &dateRFC3339Z,
						SubTotal: &positive,
					},
				)
			}

			date := Transform200601021504(resource.Updated)
			d.PatientsSummary.Date = &date
		}
	}

	return nil
}

func (d *DataSet) SetInspectionsSummary(webapi *WebApi) error {
	for _, resource := range webapi.Result.Resources {
		if regexp.MustCompile(`^\d{8}_検査実施状況.csv$`).MatchString(resource.Name) {
			// Get csv
			inspectionsStatus := NewInspectionsStatus()
			if err := inspectionsStatus.Get(resource.URL); err != nil {
				return err
			}

			// Set Dataset
			inAomori := []*int{}
			others := []*int{}
			labels := []*string{}
			zero := 0
			for _, v := range *inspectionsStatus {
				date0102, err := Transform0102(v.Date)
				if err != nil {
					return err
				}

				number := v.Number

				inAomori = append(inAomori, &number)
				others = append(others, &zero)
				labels = append(labels, &date0102)
			}

			d.InspectionsSummary.Labels = labels
			d.InspectionsSummary.Data = InspectionsSummaryData{
				InAomori: inAomori,
				Others:   others,
			}

			date := Transform200601021504(resource.Updated)
			d.InspectionsSummary.Date = &date
		}
	}

	return nil
}

func (d *DataSet) SetInspectionPersons(webapi *WebApi) error {
	for _, resource := range webapi.Result.Resources {
		if regexp.MustCompile(`^\d{8}_検査実施状況.csv$`).MatchString(resource.Name) {
			// Get csv
			inspectionsStatus := NewInspectionsStatus()
			if err := inspectionsStatus.Get(resource.URL); err != nil {
				return err
			}

			// Set Dataset
			count := []*int{}
			labels := []*string{}
			for _, v := range *inspectionsStatus {
				dateRFC3339Z, err := TransformRFC3339Z(v.Date)
				if err != nil {
					return err
				}

				number := v.Number

				labels = append(labels, &dateRFC3339Z)
				count = append(count, &number)
			}

			d.InspectionPersons.Labels = labels
			_label := "検査実施人数"
			d.InspectionPersons.Datasets = append(
				d.InspectionPersons.Datasets,
				InspectionPersonsDataset{
					Label: &_label,
					Data:  count,
				},
			)

			date := Transform200601021504(resource.Updated)
			d.InspectionPersons.Date = &date
		}
	}

	return nil
}

func (d *DataSet) SetInspectionStatusSummary(webapi *WebApi) error {
	for _, resource := range webapi.Result.Resources {
		if regexp.MustCompile(`^\d{8}_検査実施状況.csv$`).MatchString(resource.Name) {
			// Get csv
			inspectionsStatus := NewInspectionsStatus()
			if err := inspectionsStatus.Get(resource.URL); err != nil {
				return err
			}

			// Set Dataset
			total := 94
			totalInAomori := 0

			for _, v := range *inspectionsStatus {
				number := v.Number
				positive := v.Positive

				total = total + number
				totalInAomori = totalInAomori + positive

			}

			attr := "検査実施人数(累計)"
			d.InspectionStatusSummary.Attr = &attr
			d.InspectionStatusSummary.Value = &total

			var inspectionStatusSummaryChildren InspectionStatusSummaryChildren

			attr2 := "県内発生"
			inspectionStatusSummaryChildren.Children = append(
				inspectionStatusSummaryChildren.Children,
				InspectionStatusSummaryChildrenChildren{
					Attr:  &attr2,
					Value: &totalInAomori,
				},
			)

			attr3 := "その他(チャーター便・クルーズ船等)"
			zero := 0
			inspectionStatusSummaryChildren.Children = append(
				inspectionStatusSummaryChildren.Children,
				InspectionStatusSummaryChildrenChildren{
					Attr:  &attr3,
					Value: &zero,
				},
			)

			attr4 := "検査実施件数(累計)"
			d.InspectionStatusSummary.Children = append(
				d.InspectionStatusSummary.Children,
				InspectionStatusSummaryChildren{
					Attr:     &attr4,
					Value:    &total,
					Children: inspectionStatusSummaryChildren.Children,
				},
			)

			date := Transform200601021504(resource.Updated)
			d.InspectionStatusSummary.Date = &date
		}
	}

	return nil
}

func (d *DataSet) SetContacts(webapi *WebApi) error {
	for _, resource := range webapi.Result.Resources {
		if regexp.MustCompile(`^\d{8}_相談件数（コールセンター）.csv$`).MatchString(resource.Name) {

			// Get csv
			numberOfInquiriesCallCenter := NewNumberOfInquiriesCallCenter()
			if err := numberOfInquiriesCallCenter.Get(resource.URL); err != nil {
				return err
			}

			// Set Dataset
			for _, v := range *numberOfInquiriesCallCenter {
				if v.Code == "20001" {
					dateRFC3339Z, err := TransformRFC3339Z(v.Date)
					if err != nil {
						return err
					}

					date20060102, err := Transform20060102(v.Date)
					if err != nil {
						return err
					}

					date0102, err := Transform0102(v.Date)
					if err != nil {
						return err
					}

					weekday, err := TransformWeekday(v.Date)
					if err != nil {
						return err
					}

					subTotal := v.Count

					d.Contacts.Data = append(
						d.Contacts.Data,
						ContactData{
							Day:       &dateRFC3339Z,
							Date:      &date20060102,
							ShortDate: &date0102,
							W:         &weekday,
							SubTotal:  &subTotal,
						},
					)
				}
			}

			date := Transform200601021504(resource.Updated)
			d.Contacts.Date = &date
		}
	}

	return nil
}

func (d *DataSet) SetQuerents(webapi *WebApi) error {
	for _, resource := range webapi.Result.Resources {
		if regexp.MustCompile(`^\d{8}_相談件数（帰国者・接触者相談）.csv$`).MatchString(resource.Name) {

			// Get csv
			numberOfInquiriesNearCorona := NewNumberOfInquiriesNearCorona()
			if err := numberOfInquiriesNearCorona.Get(resource.URL); err != nil {
				return err
			}

			// Set Dataset
			for _, v := range *numberOfInquiriesNearCorona {
				dateRFC3339Z, err := TransformRFC3339Z(v.Date)
				if err != nil {
					return err
				}

				date20060102, err := Transform20060102(v.Date)
				if err != nil {
					return err
				}

				date0102, err := Transform0102(v.Date)
				if err != nil {
					return err
				}

				weekday, err := TransformWeekday(v.Date)
				if err != nil {
					return err
				}

				subTotal := v.Count

				d.Querents.Data = append(
					d.Querents.Data,
					QuerentData{
						Day:       &dateRFC3339Z,
						Date:      &date20060102,
						ShortDate: &date0102,
						W:         &weekday,
						SubTotal:  &subTotal,
					},
				)
			}

			date := Transform200601021504(resource.Updated)
			d.Querents.Date = &date
		}
	}

	return nil
}
