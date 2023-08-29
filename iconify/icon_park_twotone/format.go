package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Format(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFormat0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M8 24h32v18H8z"/><path fill="#555" stroke-linejoin="round" d="M4 13h14V6h12v7h14v11H4V13Z"/><path d="M16 32v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFormat0)"/>`),
		g.Group(children),
	)
}