package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nintendo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 .803h9.469L22.598 22V.803h9.401v30.395h-9.385L9.411 10.001v21.197H-.001V.803z"/>`),
		g.Group(children),
	)
}