package ant_design

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentageOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m855.7 210.8l-42.4-42.4a8.03 8.03 0 0 0-11.3 0L168.3 801.9a8.03 8.03 0 0 0 0 11.3l42.4 42.4c3.1 3.1 8.2 3.1 11.3 0L855.6 222c3.2-3 3.2-8.1.1-11.2zM304 448c79.4 0 144-64.6 144-144s-64.6-144-144-144s-144 64.6-144 144s64.6 144 144 144zm0-216c39.7 0 72 32.3 72 72s-32.3 72-72 72s-72-32.3-72-72s32.3-72 72-72zm416 344c-79.4 0-144 64.6-144 144s64.6 144 144 144s144-64.6 144-144s-64.6-144-144-144zm0 216c-39.7 0-72-32.3-72-72s32.3-72 72-72s72 32.3 72 72s-32.3 72-72 72z"/>`),
		g.Group(children),
	)
}