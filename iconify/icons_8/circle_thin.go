package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.38 4 4 9.38 4 16s5.38 12 12 12s12-5.38 12-12S22.62 4 16 4zm0 1c6.08 0 11 4.92 11 11s-4.92 11-11 11S5 22.08 5 16S9.92 5 16 5z"/>`),
		g.Group(children),
	)
}