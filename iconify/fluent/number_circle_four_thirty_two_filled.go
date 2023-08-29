package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 16C2 8.268 8.268 2 16 2s14 6.268 14 14s-6.268 14-14 14S2 23.732 2 16Zm16.998-6.179c0-1.385-1.797-1.929-2.565-.776l-5.954 8.934A1.3 1.3 0 0 0 11.56 20h5.437v2a1 1 0 1 0 2 0v-2h1a1 1 0 1 0 0-2h-1V9.821Zm-2 1.982V18h-4.13l4.13-6.197Z"/>`),
		g.Group(children),
	)
}