package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DateComesBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 8H9C7.34315 8 6 9.34315 6 11V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V25V18"/><path d="M32.2427 12.4854L36.4853 8.24271L32.2427 4.00007"/><path d="M24 17V34"/><path d="M32 24V34"/><path d="M16 24V34"/></g>`),
		g.Group(children),
	)
}