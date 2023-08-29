package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5Zm-5.93 10.64a1.42 1.42 0 0 1 .82.23l15.75 10.2a1.44 1.44 0 0 1 0 2.41l-15.75 10.2A1.44 1.44 0 0 1 16.67 35V14.57a1.43 1.43 0 0 1 1.4-1.43Z"/>`),
		g.Group(children),
	)
}