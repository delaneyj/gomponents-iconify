package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c3.9 0 7 3.1 7 7c0 5.2-7 13-7 13S5 14.2 5 9c0-3.9 3.1-7 7-7m-1 4v6h2V6h-2m0 8v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}