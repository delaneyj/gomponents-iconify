package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTToolkit0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 25v13a3 3 0 0 0 3 3h26a3 3 0 0 0 3-3V25"/><path fill="#555" d="M5 15a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-8Z"/><path stroke-linecap="round" d="M31 13V9a2 2 0 0 0-2-2H19a2 2 0 0 0-2 2v4m-2 10v6m18-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTToolkit0)"/>`),
		g.Group(children),
	)
}