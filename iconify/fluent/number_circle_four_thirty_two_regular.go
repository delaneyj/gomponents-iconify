package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 16C4 9.373 9.373 4 16 4s12 5.373 12 12s-5.373 12-12 12S4 22.627 4 16ZM16 2C8.268 2 2 8.268 2 16s6.268 14 14 14s14-6.268 14-14S23.732 2 16 2Zm2.998 7.821c0-1.385-1.797-1.929-2.565-.776l-5.954 8.934A1.3 1.3 0 0 0 11.56 20h5.437v2a1 1 0 1 0 2 0v-2h1a1 1 0 1 0 0-2h-1V9.821Zm-2 1.982V18h-4.13l4.13-6.197Z"/>`),
		g.Group(children),
	)
}