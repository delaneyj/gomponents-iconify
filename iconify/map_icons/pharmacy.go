package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pharmacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M39.352 17H10.4C5.23 17 1 20.64 1 25.5c0 4.86 4.23 8.5 9.4 8.5h28.952c5.168 0 9.398-3.641 9.398-8.5c0-4.86-4.23-8.5-9.398-8.5zm.003 15H25V19h14.355c3.812 0 6.934 2.915 6.934 6.5S43.168 32 39.355 32z"/>`),
		g.Group(children),
	)
}