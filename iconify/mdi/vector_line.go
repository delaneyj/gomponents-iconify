package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 3v4.59L7.59 15H3v6h6v-4.58L16.42 9H21V3m-4 2h2v2h-2M5 17h2v2H5"/>`),
		g.Group(children),
	)
}