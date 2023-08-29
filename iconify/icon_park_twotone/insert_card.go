package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInsertCard0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M10 18H4V6h40v12h-6"/><path fill="#555" d="M12 12L4 41h40l-8-29H12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInsertCard0)"/>`),
		g.Group(children),
	)
}