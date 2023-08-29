package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M27 34h-6v10h6V34Zm4 0H17l-3-8h20l-3 8ZM21 14c2-4.61 6-10 6-10s6.36 11.49 6.36 15.23S30.74 26 27 26h-8c-2.64 0-4.36-2.14-4.36-4.78S19 10 19 10s.96 2.01 2 4Z"/>`),
		g.Group(children),
	)
}