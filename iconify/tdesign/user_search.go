package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM17.75 15a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5ZM13 17.75a4.75 4.75 0 1 1 8.74 2.578l1.674 1.671l-1.413 1.415l-1.675-1.673A4.75 4.75 0 0 1 13 17.75ZM8 16a4 4 0 0 0-4 4h7.55v2H2v-2a6 6 0 0 1 6-6h3.5v2H8Z"/>`),
		g.Group(children),
	)
}