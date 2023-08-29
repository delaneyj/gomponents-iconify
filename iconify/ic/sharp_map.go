package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 5.1L9 3L3 5.02v16.2l6-2.33l6 2.1l6-2.02V2.77L15 5.1zm0 13.79l-6-2.11V5.11l6 2.11v11.67z"/>`),
		g.Group(children),
	)
}