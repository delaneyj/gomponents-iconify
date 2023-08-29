package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPlusLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M254 136a6 6 0 0 1-6 6h-18v18a6 6 0 0 1-12 0v-18h-18a6 6 0 0 1 0-12h18v-18a6 6 0 0 1 12 0v18h18a6 6 0 0 1 6 6Zm-57.41 60.14a6 6 0 1 1-9.18 7.72C166.9 179.45 138.69 166 108 166s-58.89 13.45-79.41 37.86a6 6 0 0 1-9.18-7.72C35.14 177.41 55 164.48 77 158.25a66 66 0 1 1 62 0c22 6.23 41.86 19.16 57.59 37.89ZM108 154a54 54 0 1 0-54-54a54.06 54.06 0 0 0 54 54Z"/>`),
		g.Group(children),
	)
}