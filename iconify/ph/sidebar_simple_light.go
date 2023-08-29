package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SidebarSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 42H40a14 14 0 0 0-14 14v144a14 14 0 0 0 14 14h176a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14ZM38 200V56a2 2 0 0 1 2-2h42v148H40a2 2 0 0 1-2-2Zm180 0a2 2 0 0 1-2 2H94V54h122a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}