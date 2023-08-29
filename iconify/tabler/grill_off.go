package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrillOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8H5a6 6 0 0 0 6 6h2c.315 0 .624-.024.926-.071m2.786-1.214a5.99 5.99 0 0 0 2.284-4.49V8h-7m6.831 10.815a2 2 0 1 1-2.663-2.633M9 14l-3 6m9-2H7m8-13V4m-3 1V4M9 5V4M3 3l18 18"/>`),
		g.Group(children),
	)
}