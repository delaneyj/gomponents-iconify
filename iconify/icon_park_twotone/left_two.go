package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLeftTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 5L6 24l18 19V31h18V17H24V5Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLeftTwo0)"/>`),
		g.Group(children),
	)
}