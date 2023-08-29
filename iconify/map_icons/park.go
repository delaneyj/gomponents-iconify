package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Park(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M17.1 26h14.87l-4.893-13h-5.073L17.1 26zM9 13V9h32v4h-7.871l4.905 13H50v4H39.863l5.729 14h-5.559l-5.728-14H15.782L10.08 44H4.52l5.697-14H0v-4h12.042l4.901-13H9z"/>`),
		g.Group(children),
	)
}