package cmd

import (
	"log"

	"asdrome.com/csv-to-pb/pocketbase"
	"github.com/spf13/cobra"
)

var collection string
var endpoint string

// rootCmd represents the base command of the CLI application.
var rootCmd = &cobra.Command{
	Use:   "upload",
	Short: "A tool to load CSV data into PocketBase",
	Run: func(cmd *cobra.Command, args []string) {
		// Crear una instancia de PocketBase
		pb := pocketbase.NewPocketBase(endpoint).Collection(collection)

		// Llamar al método LoadFromCSV en la instancia de PocketBase
		err := pb.LoadFromCSV("./" + collection + ".csv")
		if err != nil {
			// Manejar el error si es necesario
			log.Fatalln(err)
		}
	},
}

// Execute runs the root command of the CLI application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

// init initializes the command line flags.
func init() {
	rootCmd.Flags().StringVarP(&collection, "collection", "c", "", "Nombre de la colección")
	rootCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "Endpoint de PocketBase")
	rootCmd.MarkFlagRequired("collection")
	rootCmd.MarkFlagRequired("endpoint")
}
