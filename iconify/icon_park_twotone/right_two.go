package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRightTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 43l18-19L24 5v12H6v14h18v12Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRightTwo0)"/>`),
		g.Group(children),
	)
}