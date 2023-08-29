package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navicon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M10 12h30v4H10zm0 10h30v4H10zm0 10h30v4H10z"/>`),
		g.Group(children),
	)
}