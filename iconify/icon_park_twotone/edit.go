package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEdit0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M7 42h36"/><path fill="#555" d="M11 26.72V34h7.317L39 13.308L31.695 6L11 26.72Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEdit0)"/>`),
		g.Group(children),
	)
}