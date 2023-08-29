package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMark0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M11 6v36"/><path fill="#555" d="M11 9h14l7 3h7a2 2 0 0 1 2 2v17a2 2 0 0 1-2 2h-7l-7-3H11V9Z"/><path stroke-linecap="round" d="M7 42h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMark0)"/>`),
		g.Group(children),
	)
}