package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSort0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m24 42l-9-13h18l-9 13Zm0-36l-9 13h18L24 6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSort0)"/>`),
		g.Group(children),
	)
}