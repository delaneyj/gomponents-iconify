package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.379 4 4 9.379 4 16s5.379 12 12 12s12-5.379 12-12S22.621 4 16 4zm0 1c6.082 0 11 4.918 11 11s-4.918 11-11 11S5 22.082 5 16S9.918 5 16 5z"/>`),
		g.Group(children),
	)
}