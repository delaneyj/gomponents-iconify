package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M0 2v2h26V2H0zm0 5v2h26V7H0zm0 5v2h26v-2H0zm0 5v2h26v-2H0zm0 5v2h26v-2H0z"/>`),
		g.Group(children),
	)
}