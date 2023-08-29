package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendStatic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.847 5.142a1 1 0 0 1 0 1.715l-6.665 4a1 1 0 0 1-1.515-.858V2.001a1 1 0 0 1 1.515-.858l6.665 4z"/>`),
		g.Group(children),
	)
}