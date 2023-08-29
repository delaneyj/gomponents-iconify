package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Combination(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.148l6.16 9.602H5.84L12 1.148Zm0 3.704L9.5 8.75h5L12 4.852ZM2 13h9v9H2v-9Zm2 2v5h5v-5H4Zm13.5 0a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM13 17.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Z"/>`),
		g.Group(children),
	)
}