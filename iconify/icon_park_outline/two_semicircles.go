package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoSemicircles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 25c0-11.046-8.954-20-20-20S4 13.954 4 25h40Zm-30 7c0 5.523 4.477 10 10 10s10-4.477 10-10H14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}