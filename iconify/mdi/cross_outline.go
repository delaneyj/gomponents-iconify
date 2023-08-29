package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 6V1h-7v5H3v7h5.5v10h7V13H21V6h-5.5m3.5 5h-5.5v10h-3V11H5V8h5.5V3h3v5H19v3Z"/>`),
		g.Group(children),
	)
}