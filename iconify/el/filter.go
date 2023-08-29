package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m0 0l415.869 415.869V1200l368.262-301.318V415.869L1200 0H0z"/>`),
		g.Group(children),
	)
}