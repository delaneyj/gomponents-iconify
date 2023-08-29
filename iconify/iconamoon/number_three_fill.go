package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 4a1 1 0 0 1 1-1h8a1 1 0 0 1 .768 1.64l-3.884 4.662A6 6 0 1 1 7.4 19.8a1 1 0 1 1 1.2-1.6A4 4 0 1 0 11 11a1 1 0 0 1-.768-1.64L13.865 5H8a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}