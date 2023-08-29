package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustContrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zm0 14a6 6 0 1 1 0-12a6 6 0 0 1 0 12zm0-11v10c2.757 0 5-2.243 5-5s-2.243-5-5-5z"/>`),
		g.Group(children),
	)
}