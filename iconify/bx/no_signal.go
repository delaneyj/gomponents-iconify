package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSignal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16h2v5H3zm4-3h2v8H7zM21 3h-2v14.59l-2-2V7h-2v6.59l-2-2V10h-1.59l-7.7-7.71l-1.42 1.42l18 18l1.42-1.42l-.71-.7V3zm-6 18h1.88L15 19.12V21zm-4 0h2v-3.88l-2-2V21z"/>`),
		g.Group(children),
	)
}