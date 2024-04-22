package main

import (
	"context"
	"flag"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"

	"github.com/oracle/oci-go-sdk/budget"
	"github.com/oracle/oci-go-sdk/common"
)

type BudgetPlugin struct {
	Prefix string
	Budget string
}

func (b BudgetPlugin) FetchMetrics() (map[string]float64, error) {
	c, err := budget.NewBudgetClientWithConfigurationProvider(common.DefaultConfigProvider())

	if err != nil {
		return nil, err
	}

	res, err := c.GetBudget(context.Background(), budget.GetBudgetRequest{
		BudgetId: &b.Budget,
	})

	if err != nil {
		return nil, err
	}

	return map[string]float64{
		"limit":      float64(*res.Budget.Amount),
		"actual":     float64(*res.Budget.ActualSpend),
		"forecasted": float64(*res.Budget.ForecastedSpend),
	}, nil
}

func (b BudgetPlugin) GraphDefinition() map[string]mp.Graphs {
	prefix := strings.Title(b.MetricKeyPrefix())
	c, err := budget.NewBudgetClientWithConfigurationProvider(common.DefaultConfigProvider())

	if err != nil {
		return nil
	}

	res, err := c.GetBudget(context.Background(), budget.GetBudgetRequest{
		BudgetId: &b.Budget,
	})

	if err != nil {
		return nil
	}

	return map[string]mp.Graphs{
		*res.DisplayName: {
			Label: prefix,
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "limit", Label: "Limit"},
				{Name: "actual", Label: "Actual"},
				{Name: "forecasted", Label: "Forecasted"},
			},
		},
	}
}

func (b BudgetPlugin) MetricKeyPrefix() string {
	if b.Prefix == "" {
		b.Prefix = "budget"
	}
	return b.Prefix
}

func main() {
	// mackerel
	optPrefix := flag.String("metric-key-prefix", "budget", "Metric key prefix")
	optTempfile := flag.String("tempfile", "", "Temp file name")

	// oci
	budget := flag.String("budget", "budget", "Budget OCID")

	flag.Parse()

	b := BudgetPlugin{
		Prefix: *optPrefix,
		Budget: *budget,
	}

	helper := mp.NewMackerelPlugin(b)
	helper.Tempfile = *optTempfile

	helper.Run()
}
