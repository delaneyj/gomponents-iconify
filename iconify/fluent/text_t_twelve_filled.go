package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2.75A.75.75 0 0 1 2.75 2h6a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0V3.5H6.5v5h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5H5v-5H3.5v.75a.75.75 0 0 1-1.5 0v-1.5Z"/>`),
		g.Group(children),
	)
}