package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesAnchorAnchorMarinaHarborPort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 8H1.5a5.5 5.5 0 0 0 11 0H11"/><circle cx="7" cy="2" r="1.5"/><path d="M7 3.5v10m-1.5-7h3"/></g>`),
		g.Group(children),
	)
}