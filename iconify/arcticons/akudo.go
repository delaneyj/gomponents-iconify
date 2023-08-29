package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Akudo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 3.5L5.273 20.473h37.454L24 3.5zm18.727 16.973L24 44.5L5.273 20.473"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 11.496l-9.905 8.977L24 34.58l9.905-14.107L24 11.496zM24 3.5v41"/>`),
		g.Group(children),
	)
}