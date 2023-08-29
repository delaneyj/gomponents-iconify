package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseCircleB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.44 25.946a5.56 5.56 0 0 1 5.56-5.56h0a5.56 5.56 0 0 1 5.56 5.56v3.615A5.56 5.56 0 0 1 24 35.12h0a5.56 5.56 0 0 1-5.56-5.56m-.001 5.561V12.879"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}