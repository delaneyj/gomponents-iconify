package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 12.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm11.5 2a.5.5 0 0 0 0-1H19v-3h2.5a.5.5 0 0 0 0-1H19A2.502 2.502 0 0 0 16.5 5V2.5a.5.5 0 0 0-1 0V5h-3V2.5a.5.5 0 0 0-1 0V5h-3V2.5a.5.5 0 0 0-1 0V5A2.502 2.502 0 0 0 5 7.5H2.5a.5.5 0 0 0 0 1H5v3H2.5a.5.5 0 0 0 0 1H5v3H2.5a.5.5 0 0 0 0 1H5A2.502 2.502 0 0 0 7.5 19v2.5a.5.5 0 1 0 1 0V19h3v2.5a.5.5 0 1 0 1 0V19h3v2.5a.5.5 0 1 0 1 0V19a2.502 2.502 0 0 0 2.5-2.5h2.5a.5.5 0 0 0 0-1H19v-3h2.5zm-3.5 4a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 6 16.5v-9A1.5 1.5 0 0 1 7.5 6h9A1.5 1.5 0 0 1 18 7.5v9zm-4-8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1z"/>`),
		g.Group(children),
	)
}