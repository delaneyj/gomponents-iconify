package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NpmVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M3 3v18h18V3H3m3 3h12v12h-3V9h-3v9H6V6z" fill="currentColor"/>`),
		g.Group(children),
	)
}