package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignTopOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M17.25 23.25h6V3.75h-6zm1.5-18h3v16.5h-3z" fill="currentColor"/><path d="M9 16.125h6V3.75H9zM10.5 5.25h3v9.375h-3z" fill="currentColor"/><path d="M.75 23.25h6V3.75h-6zm1.5-18h3v16.5h-3z" fill="currentColor"/><path d="M.75.75h22.5v1.5H.75V.75z" fill="currentColor"/>`),
		g.Group(children),
	)
}