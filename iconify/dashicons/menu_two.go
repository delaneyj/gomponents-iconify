package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M15 7V5H5v2h10zm0 4V9H5v2h10zm0 4v-2H5v2h10z" fill="currentColor"/>`),
		g.Group(children),
	)
}