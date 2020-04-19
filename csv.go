package datasetgen

import (
	"io/ioutil"
	"net/http"

	"github.com/gocarina/gocsv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type (
	PatientStatus struct {
		Number       int    `csv:"ＮＯ"`
		Code         string `csv:"全国地方公共団体コード"`
		Prefecture   string `csv:"都道府県名"`
		HealthCenter string `csv:"保健所管内"`
		PublishDate  string `csv:"公表_年月日"`
		FixedDate    string `csv:"判明_年月日"`
		Address      string `csv:"居住地"`
		Age          string `csv:"年代"`
		Sex          string `csv:"性別"`
	}

	PatientsStatus []PatientStatus

	InspectionStatus struct {
		Date     string `csv:"検査日時"`
		Number   int    `csv:"実施数"`
		Positive int    `csv:"陽性数"`
		Negative int    `csv:"陰性数"`
	}

	InspectionsStatus []InspectionStatus

	NumberOfInquiryCallCenter struct {
		Code       string `csv:"全国地方公共団体コード"`
		Prefecture string `csv:"都道府県名"`
		City       string `csv:"市区町村名"`
		Date       string `csv:"受付_年月日"`
		Count      int    `csv:"相談件数(対応)"`
	}

	NumberOfInquiriesCallCenter []NumberOfInquiryCallCenter

	NumberOfInquiryNearCorona struct {
		Code       string `csv:"全国地方公共団体コード"`
		Prefecture string `csv:"都道府県名"`
		Date       string `csv:"受付_年月日"`
		Count      int    `csv:"相談件数"`
	}

	NumberOfInquiriesNearCorona []NumberOfInquiryNearCorona
)

func NewPatientsStatus() *PatientsStatus {
	return new(PatientsStatus)
}

func (p *PatientsStatus) Get(url string) error {
	csv, err := getCsv(url)
	if err != nil {
		return err
	}

	if err := gocsv.UnmarshalBytes(csv, p); err != nil {
		return err
	}

	return nil
}

func NewInspectionsStatus() *InspectionsStatus {
	return new(InspectionsStatus)
}

func (i *InspectionsStatus) Get(url string) error {
	csv, err := getCsv(url)
	if err != nil {
		return err
	}

	if err := gocsv.UnmarshalBytes(csv, i); err != nil {
		return err
	}

	return nil
}

func NewNumberOfInquiriesCallCenter() *NumberOfInquiriesCallCenter {
	return new(NumberOfInquiriesCallCenter)
}

func (n *NumberOfInquiriesCallCenter) Get(url string) error {
	csv, err := getCsv(url)
	if err != nil {
		return err
	}

	if err := gocsv.UnmarshalBytes(csv, n); err != nil {
		return err
	}

	return nil
}

func NewNumberOfInquiriesNearCorona() *NumberOfInquiriesNearCorona {
	return new(NumberOfInquiriesNearCorona)
}

func (n *NumberOfInquiriesNearCorona) Get(url string) error {
	csv, err := getCsv(url)
	if err != nil {
		return err
	}

	if err := gocsv.UnmarshalBytes(csv, n); err != nil {
		return err
	}

	return nil
}

func getCsv(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawReader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())
	csvData, err := ioutil.ReadAll(rawReader)
	return csvData, err
}
