package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesCartCartTravelLocalBuckboardRetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.59 10H4V3a1 1 0 0 1 1-1h4.5a4 4 0 0 1 4 4v3a1 1 0 0 1-1 1h-2.09M4 10H.5V8"/><circle cx="9" cy="10.5" r="1.5"/><path d="M7.5 5h3"/></g>`),
		g.Group(children),
	)
}