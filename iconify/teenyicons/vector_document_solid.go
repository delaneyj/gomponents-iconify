package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorDocumentSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9 10v1H6v-1H5V7h1V6h3v1h1v3H9ZM4 5v1h1V5H4Zm6 0v1h1V5h-1Zm-6 7v-1h1v1H4Zm6-1v1h1v-1h-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM3 4h3v1h3V4h3v3h-1v3h1v3H9v-1H6v1H3v-3h1V7H3V4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}