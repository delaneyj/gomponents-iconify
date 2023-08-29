package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseCircleY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M29.56 22.054v7.507A5.56 5.56 0 0 1 24 35.12h0a5.543 5.543 0 0 1-3.932-1.629"/><path d="M29.56 12.879v9.175a5.56 5.56 0 0 1-5.56 5.56h0a5.56 5.56 0 0 1-5.56-5.56v-9.175"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}