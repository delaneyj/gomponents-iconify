package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2H5a1 1 0 1 1 0-2h8a1 1 0 0 1 1 1v8a1 1 0 0 1-2 0V2zM8 6H1a1 1 0 1 1 0-2h8a1 1 0 0 1 1 1v8a1 1 0 0 1-2 0V6z"/>`),
		g.Group(children),
	)
}