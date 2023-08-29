package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHomeTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 44V20L24 4L4 20v24h12V26h16v18h12Z"/><path stroke-linecap="round" d="M24 44V34"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHomeTwo0)"/>`),
		g.Group(children),
	)
}