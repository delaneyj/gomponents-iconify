package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Template(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 12h5a1 1 0 0 0 1-1V6H6v6Zm-2 0V6H2v5a1 1 0 0 0 1 1h1Zm8-9a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v1h10V3ZM3 0h8a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}