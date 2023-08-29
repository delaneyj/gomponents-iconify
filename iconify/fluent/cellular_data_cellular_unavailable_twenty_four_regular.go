package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularDataCellularUnavailableTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M3.75 18.004a.75.75 0 0 1 .743.648l.007.102v1.5a.75.75 0 0 1-1.493.101L3 20.254v-1.5a.75.75 0 0 1 .75-.75zm8-6a.75.75 0 0 1 .743.648l.007.102v7.5a.75.75 0 0 1-1.493.101L11 20.254v-7.5a.75.75 0 0 1 .75-.75zm4-3a.75.75 0 0 1 .743.648l.007.102v10.5a.75.75 0 0 1-1.493.101L15 20.254v-10.5a.75.75 0 0 1 .75-.75zm4 10.5a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zM7.75 15a.75.75 0 0 1 .743.648l.007.102v4.499a.75.75 0 0 1-1.493.101L7 20.25v-4.5a.75.75 0 0 1 .75-.75zm11.984-8.996a.744.744 0 0 1 .74.641l.009.102l.017 10.885a.746.746 0 0 1-.734.758a.744.744 0 0 1-.74-.642l-.009-.101L19 6.76a.746.746 0 0 1 .734-.757z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}