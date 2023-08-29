package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M161.98 397.63L0 256l161.98-141.63l27.65 31.61L64 256l125.63 110.02l-27.65 31.61zm188.04 0l-27.65-31.61L448 256L322.37 145.98l27.65-31.61L512 256L350.02 397.63z"/>`),
		g.Group(children),
	)
}