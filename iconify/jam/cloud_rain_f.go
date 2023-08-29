package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudRainF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 0a7 7 0 0 1 0 14H5a5 5 0 1 1 1.561-9.751A7.002 7.002 0 0 1 13 0z"/><rect width="2" height="4" x="5" y="15" rx="1"/><rect width="2" height="5" x="9" y="15" rx="1"/><rect width="2" height="3" x="13" y="15" rx="1"/></g>`),
		g.Group(children),
	)
}