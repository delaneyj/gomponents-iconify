package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpElectricalServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 13h3v2h-3zm-6-1v2h-2v4h2v2h5v-8z"/><path fill="currentColor" d="M5 11h7V4H4v2h6v3H3v8h6v-2H5zm13 6h3v2h-3z"/>`),
		g.Group(children),
	)
}