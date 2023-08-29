package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 12v26h26M18 10h18v17"/><path d="m44 14l-8 6.75v6.5L44 34V14Z" clip-rule="evenodd"/><path d="M44 44L4 4"/></g>`),
		g.Group(children),
	)
}