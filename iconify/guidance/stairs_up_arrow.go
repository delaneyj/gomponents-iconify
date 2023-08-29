package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StairsUpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 24v-.149a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12V6.5H24M10.5.5l-10 10m10-10c-.398.398-1.133.654-1.79.812c-.878.212-1.79.25-2.687.149c-.697-.08-1.464-.235-1.773-.544M10.5.5c-.398.398-.654 1.133-.812 1.79a7.745 7.745 0 0 0-.149 2.687c.08.697.235 1.464.545 1.773"/>`),
		g.Group(children),
	)
}