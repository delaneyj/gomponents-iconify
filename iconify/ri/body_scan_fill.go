package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BodyScanFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16v4h4v2H2v-6h2Zm18 0v6h-6v-2h4v-4h2ZM7.5 7a4.5 4.5 0 0 0 9 0h2a6.5 6.5 0 0 1-3.499 5.767L15 19H9v-6.232A6.5 6.5 0 0 1 5.5 7h2ZM12 5a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5ZM8 2v2l-4-.001V8H2V2h6Zm14 0v6h-2V4h-4V2h6Z"/>`),
		g.Group(children),
	)
}