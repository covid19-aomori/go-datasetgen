package datasetgen

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	// Request webapi
	webapi := NewWebApi()
	if err := webapi.Get(); err != nil {
		return err
	}

	// Initialize dataset
	dataset := NewDataSet()
	dataset.SetLastUpdateDate()

	// goroutine
	wg := sync.WaitGroup{}

	// .patients
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := dataset.SetPatients(webapi); err != nil {
			panic(err)
		}
	}()

	// .patients_summary
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := dataset.SetPatientsSummary(webapi); err != nil {
			panic(err)
		}
	}()

	// .inspections_summary
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := dataset.SetInspectionsSummary(webapi); err != nil {
			panic(err)
		}
	}()

	// .inspection_persons
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := dataset.SetInspectionPersons(webapi); err != nil {
			panic(err)
		}
	}()

	// .inspection_status_summary
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := dataset.SetInspectionStatusSummary(webapi); err != nil {
			panic(err)
		}
	}()

	// .contacts
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := dataset.SetContacts(webapi); err != nil {
			panic(err)
		}
	}()

	// .querents
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := dataset.SetQuerents(webapi); err != nil {
			panic(err)
		}
	}()

	// goroutine
	wg.Wait()

	// Genrate json
	datasetJson, err := dataset.GenerateJson()
	if err != nil {
		return err
	}

	fmt.Println(datasetJson)

	return nil
}
