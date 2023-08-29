package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 9.9L7 12l3 2.1V9.9m9 0L16 12l3 2.1V9.9M12 6v12l-8.5-6L12 6m9 0v12l-8.5-6L21 6Z"/>`),
		g.Group(children),
	)
}