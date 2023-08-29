package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTower0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M17 31h14v13H17zm2-14h10v14H19zm2-13h6v13h-6z"/><path stroke-linecap="round" d="M4 44h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTower0)"/>`),
		g.Group(children),
	)
}