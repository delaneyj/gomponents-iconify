package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LgThinq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="8.804" cy="36.714" r="3.304" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.196" cy="17.554" r="3.304" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.804 33.41V17.555L24 6.982l12.464 8.67"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.196 20.857v18.5H19.375V24.161h10.571v9.25"/>`),
		g.Group(children),
	)
}