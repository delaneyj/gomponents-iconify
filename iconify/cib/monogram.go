package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.875 0v31.339c0 .599-.708.891-1.104.448L1.126 0zM14.703 17.948L2.23 31.787c-.396.443-1.104.151-1.104-.448V0z"/>`),
		g.Group(children),
	)
}