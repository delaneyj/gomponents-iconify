package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13z"/><path fill="currentColor" d="M10.457 4.957L9.043 3.543L4.586 8l4.457 4.457l1.414-1.414L7.414 8z"/>`),
		g.Group(children),
	)
}