package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDisconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.739 5.819a5 5 0 0 1 6.563 6.08a2 2 0 1 1 2.234 3.312m-2.104.289H5.5a3 3 0 1 1 .1-5.998a4.988 4.988 0 0 1 1.286-2.457M4 4l13 13.071"/>`),
		g.Group(children),
	)
}