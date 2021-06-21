package view

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/GolangSriLanka/go-puso/transact/puso"
)

var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "view puso as a table",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	var s [][]string

	p := puso.Puso{}
	data, err := p.GetList()
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range data {
		s = append(s, []string{
			strconv.Itoa(int(d.ID)),
			d.Name,
			d.Color,
			fmt.Sprintf("%f", d.Weight),
			d.Owner,
			fmt.Sprintf("%d", d.Laziness),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Color", "Weight", "Owner", "Laziness"})
	table.SetBorder(false)
	table.AppendBulk(s)
	table.Render()
}
