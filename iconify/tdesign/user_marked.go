package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm8.75 6h8.5v10.294l-4.247-2.616l-4.253 2.614V13.5Zm2 2v4.715l2.254-1.385l2.246 1.383V15.5h-4.5ZM8 16a4 4 0 0 0-4 4h8.8v2H2v-2a6 6 0 0 1 6-6h4.75v2H8Z"/>`),
		g.Group(children),
	)
}