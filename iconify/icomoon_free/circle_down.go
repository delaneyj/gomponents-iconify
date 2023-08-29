package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 8A8 8 0 1 0 0 8a8 8 0 0 0 16 0zM1.5 8a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0z"/><path fill="currentColor" d="M4.957 5.543L3.543 6.957L8 11.414l4.457-4.457l-1.414-1.414L8 8.586z"/>`),
		g.Group(children),
	)
}