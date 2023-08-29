package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignCenterOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M17.25 23.25h6v-9h-6zm1.5-7.5h3v6h-3z" fill="currentColor"/><path d="M9 19.875h6V14.25H9zm1.5-4.125h3v2.625h-3z" fill="currentColor"/><path d="M.75 23.25h6v-9h-6zm1.5-7.5h3v6h-3z" fill="currentColor"/><path d="M6.75.75h-6v9h6zm-1.5 7.5h-3v-6h3z" fill="currentColor"/><path d="M15 4.5H9v5.25h6zm-1.5 3.75h-3V6h3z" fill="currentColor"/><path d="M17.25.75v9h6v-9zm4.5 7.5h-3v-6h3z" fill="currentColor"/><path d="M.75 11.25h22.5v1.5H.75v-1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}