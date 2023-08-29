package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSort0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m24 42l-9-13h18l-9 13Zm0-36l-9 13h18L24 6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSort0)"/>`),
		g.Group(children),
	)
}