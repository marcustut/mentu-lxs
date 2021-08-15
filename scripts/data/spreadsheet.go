package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SpreadsheetService struct {
	service *sheets.Service
}

func NewSpreadSheetService(clientSecretPath string) (*SpreadsheetService, error) {
	ctx := context.Background()

	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	srv, err := sheets.NewService(
		ctx,
		option.WithCredentialsFile(fmt.Sprintf("%s%s", path, clientSecretPath)),
		option.WithScopes(sheets.SpreadsheetsScope),
	)
	if err != nil {
		return nil, err
	}

	c := &SpreadsheetService{service: srv}

	return c, nil
}

func (s *SpreadsheetService) Read(spreadsheetId string, range_ string) ([][]interface{}, error) {
	values, err := s.service.Spreadsheets.Values.Get(spreadsheetId, range_).Do()

	return values.Values, err
}

func (s *SpreadsheetService) Write(
	spreadsheetId string,
	range_ string,
	values []interface{},
) error {
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, values)

	_, err := s.service.Spreadsheets.Values.Append(spreadsheetId, range_, &vr).
		ValueInputOption("RAW").
		Do()

	return err
}
