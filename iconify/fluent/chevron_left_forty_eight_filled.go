package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M32.06 8.19a1.5 1.5 0 0 1 0 2.12L18.622 23.75l13.44 13.44a1.5 1.5 0 0 1-2.122 2.12l-14.5-14.5a1.5 1.5 0 0 1 0-2.12l14.5-14.5a1.5 1.5 0 0 1 2.122 0Z"/>`),
		g.Group(children),
	)
}