package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M430.078 0v430.078H0v339.844h430.078V1200h339.844V769.922H1200V430.078H769.922V0H430.078z"/>`),
		g.Group(children),
	)
}