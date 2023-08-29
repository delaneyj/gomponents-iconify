package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 5.5A2.5 2.5 0 0 1 3.5 3h9A2.5 2.5 0 0 1 15 5.5v4a2.5 2.5 0 0 1-2.5 2.5h-9A2.5 2.5 0 0 1 1 9.5v-4ZM3.5 4A1.5 1.5 0 0 0 2 5.5v4A1.5 1.5 0 0 0 3.5 11h9A1.5 1.5 0 0 0 14 9.5v-4A1.5 1.5 0 0 0 12.5 4h-9ZM7 7.5a.5.5 0 0 1 .5-.5H12a.5.5 0 0 1 0 1H7.5a.5.5 0 0 1-.5-.5Zm-1 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}