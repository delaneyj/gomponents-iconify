package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoClipMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 3A3.25 3.25 0 0 0 2 6.25v9a3.25 3.25 0 0 0 3.25 3.25h10.5A3.25 3.25 0 0 0 19 15.25v-9A3.25 3.25 0 0 0 15.75 3H5.25ZM8 13.25V7.751a.75.75 0 0 1 1.155-.631l4.618 2.959a.5.5 0 0 1 0 .842L9.155 13.88A.75.75 0 0 1 8 13.249ZM7.75 21a3.247 3.247 0 0 1-2.74-1.5h11.24A3.75 3.75 0 0 0 20 15.75V6.01a3.247 3.247 0 0 1 1.5 2.74v7c0 2.9-2.35 5.25-5.25 5.25h-8.5Z"/>`),
		g.Group(children),
	)
}