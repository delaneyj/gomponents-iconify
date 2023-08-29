package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColorSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.46 1.303a.5.5 0 0 0-.92 0l-3 7a.5.5 0 1 0 .92.394L6.187 7h3.626l.728 1.697a.5.5 0 1 0 .919-.394l-3-7ZM9.385 6h-2.77L8 2.77L9.385 6ZM3.5 10A1.5 1.5 0 0 0 2 11.5v2A1.5 1.5 0 0 0 3.5 15h9a1.5 1.5 0 0 0 1.5-1.5v-2a1.5 1.5 0 0 0-1.5-1.5h-9ZM3 11.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5v-2Z"/>`),
		g.Group(children),
	)
}