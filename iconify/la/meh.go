package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.383 4 4 9.383 4 16s5.383 12 12 12s12-5.383 12-12S22.617 4 16 4zm0 2c5.535 0 10 4.465 10 10s-4.465 10-10 10S6 21.535 6 16S10.465 6 16 6zm-4.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm9 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zM11 20v2h10v-2z"/>`),
		g.Group(children),
	)
}