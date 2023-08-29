package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a4 4 0 0 0-4 4v6a4 4 0 0 0 8 0V9a4 4 0 0 0-4-4ZM6 9a6 6 0 1 1 12 0v6a6 6 0 0 1-12 0V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}