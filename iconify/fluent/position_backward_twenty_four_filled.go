package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PositionBackwardTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.5 6.5h-5.75a4.25 4.25 0 0 0-4.25 4.25v5.75H5.25A3.25 3.25 0 0 1 2 13.25v-8A3.25 3.25 0 0 1 5.25 2h8a3.25 3.25 0 0 1 3.25 3.25V6.5ZM22 18.75A3.25 3.25 0 0 1 18.75 22h-8a3.25 3.25 0 0 1-3.25-3.25v-8a3.25 3.25 0 0 1 3.25-3.25h8A3.25 3.25 0 0 1 22 10.75v8Zm-3.25 1.75a1.75 1.75 0 0 0 1.75-1.75v-8A1.75 1.75 0 0 0 18.75 9h-8A1.75 1.75 0 0 0 9 10.75v8c0 .966.784 1.75 1.75 1.75h8Z"/>`),
		g.Group(children),
	)
}