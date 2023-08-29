package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M649 94H0V1h649v93zm0 231H0v-92h649v92zm0 232H0v-93h649v93z"/>`),
		g.Group(children),
	)
}