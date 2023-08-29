package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M0 2v2h26V2H0zm4 5v2h18V7H4zm-4 5v2h26v-2H0zm4 5v2h18v-2H4zm-4 5v2h26v-2H0zm902 1447v2h26v-2h-26zm4 5v2h18v-2h-18zm-4 5v2h26v-2h-26zm4 5v2h18v-2h-18zm-4 5v2h26v-2h-26z"/>`),
		g.Group(children),
	)
}