package dashboard

import (
	"fmt"
	"path"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func bar(data values) components.Charter {
	Revert(data)

	names := data.barNames()

	var barData []opts.BarData
	for _, v := range data {
		barData = append(barData, opts.BarData{
			Name:  path.Join(v.Alias, v.Package, v.Name),
			Value: v.Value,
		})
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Fix commits",
			Subtitle: "Changes with 'fix' or 'bug' applied to file",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Package",
			Type: "category",
			Show: true,
			Data: names,
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				ShowMinLabel: true,
				ShowMaxLabel: true,
			},
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Show: true,
			Name: "Fixes count",
			Type: "value",
		}),
		charts.WithGridOpts(opts.Grid{
			ContainLabel: true,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "100%",
			Height: fmt.Sprintf("%dpx", 20*len(barData)),
		}),
	)

	bar.AddSeries("", barData)

	return bar
}