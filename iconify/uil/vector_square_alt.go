package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorSquareAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 7H8a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1Zm-1 8H9V9h6Zm6 3.28V5.72A2 2 0 1 0 18.28 3H5.72A2 2 0 1 0 3 5.72v12.56A2 2 0 1 0 5.72 21h12.56A2 2 0 1 0 21 18.28Zm-2 0a1.91 1.91 0 0 0-.72.72H5.72a1.91 1.91 0 0 0-.72-.72V5.72A1.91 1.91 0 0 0 5.72 5h12.56a1.91 1.91 0 0 0 .72.72Z"/>`),
		g.Group(children),
	)
}