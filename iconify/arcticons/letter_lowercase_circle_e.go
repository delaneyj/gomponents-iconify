package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseCircleE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.831 28.562A5.559 5.559 0 0 1 24 31.368h0a5.56 5.56 0 0 1-5.56-5.56v-3.615a5.56 5.56 0 0 1 5.56-5.56h0a5.56 5.56 0 0 1 5.56 5.56V24H18.44"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}