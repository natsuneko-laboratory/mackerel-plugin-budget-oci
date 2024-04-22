# mackerel-plugin-budget-oci

Budget custom metrics plugin for mackerel.io agent.

## Install

```bash
$ mkr plugin install natsuneko-laboratory/mackerel-plugin-budget-oci
```

## Setting

```toml
[plugin.metrics.budget]
command = "/path/to/mackerel-plugin-budget"
```

## Example Metrics

```bash
$ mackerel-plugin-budget-oci -budget=ocid1.budget.oc1.ap-tokyo-1.a...
budget.BudgetA.limit      100     1713804022
budget.BudgetA.actual     151.242233      1713804022
budget.BudgetA.forecasted 206.100006      1713804022
```

## License

MIT by [@6jz](https://twitter.com/6jz)
