package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 17h18m-20 9h18m-21 9h18m-1.435 8l9.74-35.47A2 2 0 0 0 36.377 5H22.214a2 2 0 0 0-1.91 1.41L9 43m8-25l4 24m14-24l4 24"/>`),
		g.Group(children),
	)
}