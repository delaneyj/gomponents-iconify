package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerfumerBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="24" x="5" y="17" rx="2"/><path d="M14 7h20v10H14zm4 18h12v8H18zm12 4h13M5 29h13M5 24v10m38-10v10"/></g>`),
		g.Group(children),
	)
}