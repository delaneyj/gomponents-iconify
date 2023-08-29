package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 6v19h10c1.102 0 2 .898 2 2h2c0-1.102.898-2 2-2h10V6H19a3.997 3.997 0 0 0-3 1.36A3.997 3.997 0 0 0 13 6zm2 2h8c1.102 0 2 .898 2 2h2c0-1.102.898-2 2-2h8v15h-8a3.997 3.997 0 0 0-3 1.36A3.997 3.997 0 0 0 13 23H5zm10 4v2h2v-2zm0 4v2h2v-2zm0 4v2h2v-2z"/>`),
		g.Group(children),
	)
}