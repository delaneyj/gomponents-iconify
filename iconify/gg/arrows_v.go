package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.757 5.04l1.415 1.415L11 4.627V10h2V4.627l1.828 1.828l1.415-1.414L12 .798L7.757 5.041Zm8.486 13.92l-1.415-1.415L13 19.373V14h-2v5.373l-1.828-1.828l-1.415 1.414L12 23.202l4.243-4.243Z"/>`),
		g.Group(children),
	)
}