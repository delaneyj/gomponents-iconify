package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Projects(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M9 15v8H1v-8h8Zm14 0v8h-8v-8h8ZM9 1v8H1V1h8Zm14 0v8h-8V1h8Z"/>`),
		g.Group(children),
	)
}