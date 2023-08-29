package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocDetail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDocDetail0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M39 4H11a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path d="M17 30h14m-14 6h7"/><path fill="#555" d="M17 12h14v10H17z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDocDetail0)"/>`),
		g.Group(children),
	)
}