package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ezmeral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 8h34v8H7V8ZM1 8h6v8H1V8Zm40 0h6v8h-6V8ZM7 16h34v6H7v-6ZM7 2h34v6H7V2ZM1 8l6-6v6H1Zm0 8l6 6v-6H1Zm46-8l-6-6v6h6Zm0 8l-6 6v-6h6Z"/>`),
		g.Group(children),
	)
}