package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkNow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.924 9.174H32.66c15.436 0 15.014 28.994-5.327 29.652v-3.668H16.225c-16.39 0-14.762-25.985-.301-25.985Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.66 24.895c5.012 6.722 15.075 6.768 20.683 0"/>`),
		g.Group(children),
	)
}