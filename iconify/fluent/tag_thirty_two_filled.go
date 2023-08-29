package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.952 4.079A4 4 0 0 1 18.684 3H25.5A3.5 3.5 0 0 1 29 6.5v6.757a4 4 0 0 1-1.172 2.829L16.005 27.909a4.25 4.25 0 0 1-6.01 0l-6.326-6.326a4.25 4.25 0 0 1 .101-6.109L15.952 4.08ZM22.5 12a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}