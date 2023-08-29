package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.5A4.5 4.5 0 0 1 7.5 3h17A4.5 4.5 0 0 1 29 7.5v17a4.5 4.5 0 0 1-4.5 4.5h-17A4.5 4.5 0 0 1 3 24.5v-17ZM7.5 5A2.5 2.5 0 0 0 5 7.5V15h10V5H7.5ZM17 5v10h10V7.5A2.5 2.5 0 0 0 24.5 5H17Zm-2 12H5v7.5A2.5 2.5 0 0 0 7.5 27H15V17Zm2 10h7.5a2.5 2.5 0 0 0 2.5-2.5V17H17v10Z"/>`),
		g.Group(children),
	)
}