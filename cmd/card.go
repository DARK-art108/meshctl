package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card <>",
	Short: "This cmd creates awesome cards",
	Long: `For example: 
	meshctl create card ritesh -n="Ritesh Yadav" -o=thanksgiving -l=fr
	meshctl create card bob --name="Bob Marley" -ocassion=birthday`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("Language")
		fmt.Println("value of flag name :" + name)
		fmt.Println("value of the flag language :" + language)
	},
}

func init() {
	createCmd.AddCommand(cardCmd)
	cardCmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear,thanksgiving, birthday, diwali, holi")
	cardCmd.PersistentFlags().StringP("Language", "l", "en", "Possible values: en, ch")
	cardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	cardCmd.MarkPersistentFlagRequired("name")
	cardCmd.MarkPersistentFlagRequired("occasion")

}
