package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 12.8V17h14V9.9c-2.5.6-4.8 1.6-7 3V10c-2.2 1.3-4.5 2.3-7 2.8zM17 3H3v7.1c2.5-.6 4.9-1.6 7-3.1v3c2.1-1.4 4.5-2.3 7-2.8V3z"/>`),
		g.Group(children),
	)
}