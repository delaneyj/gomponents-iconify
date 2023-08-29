package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 25h-5L7 15c2-.667 6-4 6-12h6c0 7.6 4 11.167 6 12l-4 10h-5Zm0 0v-6m0-6V3M9 25h14v4H9v-4Zm9-9a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}