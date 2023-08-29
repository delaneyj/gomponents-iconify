package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToyaGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.04 17.003A9.75 9.75 0 1 0 24 24m2.96 6.997A9.75 9.75 0 1 0 24 24m0 0h-9.75"/>`),
		g.Group(children),
	)
}