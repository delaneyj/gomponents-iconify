package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deliver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 18H1V3h13v14m0 1H9m-3 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm11 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM14 8h5l4 5v5h-3"/>`),
		g.Group(children),
	)
}