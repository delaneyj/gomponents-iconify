package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Manual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M14 9v8v-8Zm-4 0v8v-8ZM8 5a4 4 0 1 0 8 0a4 4 0 0 0-8 0ZM4 23h16v-3H4v3Zm3-3h10v-3H7v3Z"/>`),
		g.Group(children),
	)
}