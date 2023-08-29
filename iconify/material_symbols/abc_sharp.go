package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbcSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 15V9h5v2h-1.5v-.5h-2v3h2V13H21v2h-5Zm-6.5 0V9h4.25l.75.75v1.5l-.75.75l.75.75v1.5l-.75.75H9.5Zm1.5-3.75h2v-.75h-2v.75Zm0 2.25h2v-.75h-2v.75ZM3 15V9h5v6H6.5v-1.5h-2V15H3Zm1.5-3h2v-1.5h-2V12Z"/>`),
		g.Group(children),
	)
}