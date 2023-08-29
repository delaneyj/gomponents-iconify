package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpDownSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 14V5.4L5.4 0l1.225 1.225L5.8 5H12v2.575L9.275 14H0Zm18.6 10l-1.225-1.225L18.2 19H12v-2.575L14.725 10H24v8.6L18.6 24Z"/>`),
		g.Group(children),
	)
}