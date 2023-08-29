package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 21h-5v-8.25H21v5A3.25 3.25 0 0 1 17.75 21ZM21 11.25h-8.25V3h5A3.25 3.25 0 0 1 21 6.25v5Zm-9.75 0V3h-5A3.25 3.25 0 0 0 3 6.25v5h8.25ZM3 12.75v5A3.25 3.25 0 0 0 6.25 21h5v-8.25H3Z"/>`),
		g.Group(children),
	)
}