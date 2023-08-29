package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Text(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1v2a1 1 0 0 1-2 0V2H7v8h1a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2h1V2H2v1a1 1 0 1 1-2 0V1a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}