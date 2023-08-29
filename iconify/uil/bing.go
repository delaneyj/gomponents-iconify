package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.1 8.6l1.7 4.3l2.8 1.3L9 17.5V3.4L5 2v17.8L9 22l10-5.8v-4.5l-8.9-3.1z"/>`),
		g.Group(children),
	)
}