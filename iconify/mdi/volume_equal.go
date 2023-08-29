package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9h4l5-5v16l-5-5H3V9m11 4h8v2h-8m0-6h8v2h-8Z"/>`),
		g.Group(children),
	)
}