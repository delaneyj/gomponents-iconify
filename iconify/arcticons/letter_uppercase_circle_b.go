package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterUppercaseCircleB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.881 36.15v-24.3m-.001 0h10.177c3.348 0 6.062 2.72 6.062 6.075h0c0 3.355-2.714 6.075-6.062 6.075H15.881m-.001 0h10.177c3.348 0 6.062 2.72 6.062 6.075h0c0 3.355-2.714 6.075-6.062 6.075H15.881"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}