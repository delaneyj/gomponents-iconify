package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseCircleP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.44 22.054a5.56 5.56 0 0 0 5.56 5.56h0a5.56 5.56 0 0 0 5.56-5.56v-3.615A5.56 5.56 0 0 0 24 12.88h0a5.56 5.56 0 0 0-5.56 5.56m-.001-5.561v22.242"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}