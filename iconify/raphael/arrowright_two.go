package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowrightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.13 22.186L16.315 16L10.13 9.81l3.535-3.536L23.39 16l-9.725 9.725z"/>`),
		g.Group(children),
	)
}