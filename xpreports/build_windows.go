package xpreports

import (
	"strconv"

	"github.com/w0de/gosal/config"
	"github.com/w0de/gosal/xpreports/windows"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// buildReport creates a report using windows APIs and paths.
func buildReport(conf *config.Config) (*Report, error) {
	win32Bios, err := windows.GetWin32Bios()
	if err != nil {
		return nil, errors.Wrap(err, "get win32Bios")
	}

	CDrive, err := windows.GetCDrive()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting win32 disk")
	}

	u1 := uuid.NewV4().String()

	encodedCompressedPlist, err := windows.BuildBase64bz2Report(conf)
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting plist")
	}

	report := &Report{
		Serial:          win32Bios.SerialNumber,
		Key:             conf.Key,
		Name:            win32Bios.PSComputerName,
		DiskSize:        strconv.Itoa(CDrive.Size),
		SalVersion:      strconv.Itoa(1),
		RunUUID:         u1,
		Base64bz2Report: encodedCompressedPlist,
	}

	// fmt.Printf("%+v\n", report)
	return report, nil
}
