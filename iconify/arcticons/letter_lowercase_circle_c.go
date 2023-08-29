package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseCircleC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.195 28.567a5.559 5.559 0 0 1-4.829 2.8h0a5.56 5.56 0 0 1-5.56-5.56v-3.614a5.56 5.56 0 0 1 5.56-5.56h0a5.558 5.558 0 0 1 4.823 2.79"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}