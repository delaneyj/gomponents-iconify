package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignBottomOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M6.75.75h-6v19.875h6zm-1.5 18.375h-3V2.25h3z" fill="currentColor"/><path d="M15 7.948H9v12.677h6zm-1.5 11.177h-3V9.448h3z" fill="currentColor"/><path d="M17.25.75v19.875h6V.75zm4.5 18.375h-3V2.25h3z" fill="currentColor"/><path d="M.75 21.75h22.5v1.5H.75v-1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}