package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticallyCentered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVerticallyCentered0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M6 7h36"/><path fill="#555" stroke-linejoin="round" d="M16 16h16v16H16z"/><path d="M6 41h36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVerticallyCentered0)"/>`),
		g.Group(children),
	)
}