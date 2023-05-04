package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func run() {
	var format string
	var command string
	var objData string
	var makeEphem string
	var ephemType string
	var center string
	var startTime string
	var stopTime string
	var stepSize string
	var quantities string

	rootCmd := &cobra.Command{
		Use:   "horizons",
		Short: "Download data from the NASA Horizons system",
		Run: func(cmd *cobra.Command, args []string) {
			api := &HorizonsAPI{
				Format:     format,
				Command:    command,
				ObjData:    objData,
				MakeEphem:  makeEphem,
				EphemType:  ephemType,
				Center:     center,
				StartTime:  startTime,
				StopTime:   stopTime,
				StepSize:   stepSize,
				Quantities: quantities,
			}

			response, err := api.download()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error downloading data: %s\n", err)
				os.Exit(1)
			}

			fmt.Println(string(response))
		},
	}

	rootCmd.Flags().StringVarP(&format, "format", "f", "text", "Output format")
	rootCmd.Flags().StringVarP(&command, "command", "c", "", "Horizons command")
	rootCmd.Flags().StringVarP(&objData, "obj-data", "o", "YES", "Include object data")
	rootCmd.Flags().StringVarP(&makeEphem, "make-ephem", "e", "YES", "Generate ephemerides")
	rootCmd.Flags().StringVarP(&ephemType, "ephem-type", "t", "OBSERVER", "Ephemeris type")
	rootCmd.Flags().StringVarP(&center, "center", "r", "'500@399'", "Observation center")
	rootCmd.Flags().StringVarP(&startTime, "start-time", "s", "", "Start time (ISO 8601 format)")
	rootCmd.Flags().StringVarP(&stopTime, "stop-time", "x", "", "Stop time (ISO 8601 format)")
	rootCmd.Flags().StringVarP(&stepSize, "step-size", "p", "1 d", "Step size")
	rootCmd.Flags().StringVarP(&quantities, "quantities", "q", "", "Output quantities")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
		os.Exit(1)
	}
}
func main() {

	run()

}
