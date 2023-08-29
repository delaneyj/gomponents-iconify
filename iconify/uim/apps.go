package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="9" height="9" x="2" y="2" fill="currentColor" rx="1"/><rect width="9" height="9" x="2" y="13" fill="currentColor" opacity=".5" rx="1"/><rect width="9" height="9" x="13" y="2" fill="currentColor" opacity=".5" rx="1"/><rect width="9" height="9" x="13" y="13" fill="currentColor" opacity=".5" rx="1"/>`),
		g.Group(children),
	)
}