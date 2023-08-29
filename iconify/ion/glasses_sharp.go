package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M496 176H16v64h21.24l12.44 112h171.87L240 241.32V240a16 16 0 0 1 32 0v1.32L290.45 352h171.87l12.44-112H496Z"/>`),
		g.Group(children),
	)
}