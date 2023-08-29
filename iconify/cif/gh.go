package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#006B3F" d="M.5.5h300v200H.5"/><path fill="#FCD116" d="M.5.5h300v133.333H.5"/><path fill="#CE1126" d="M.5.5h300v66.667H.5"/><path fill="#000" d="M114.5 92.5h70.667L128.5 133.833L149.834 66.5l22.666 67.333"/></g>`),
		g.Group(children),
	)
}