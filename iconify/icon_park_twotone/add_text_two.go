package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddTextTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAddTextTwo0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 16H16m8 18V16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAddTextTwo0)"/>`),
		g.Group(children),
	)
}