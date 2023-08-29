package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m0 4.599l13-1.796v12.599H0zm14.599-2L32 0v15.197H14.599zM0 16.803h13v12.599L0 27.599zm14.599 0H32V32l-17.197-2.401z"/>`),
		g.Group(children),
	)
}