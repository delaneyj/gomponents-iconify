package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRoadSign0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 10v12h32l6-6l-6-6H6Z"/><path stroke-linecap="round" d="M23 22v22m0-40v6m-5 34h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRoadSign0)"/>`),
		g.Group(children),
	)
}