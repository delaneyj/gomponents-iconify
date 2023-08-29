package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorLineAccentSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.007 11.414a1.748 1.748 0 0 0 1.958.773l2.397-.655h.001A3.17 3.17 0 0 0 7.44 11h6.06a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-2c0-.03.003-.058.007-.086Z"/>`),
		g.Group(children),
	)
}