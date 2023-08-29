package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 13v-2h14.485l-3.242-3.243l1.414-1.414L19.314 12l-5.657 5.657l-1.414-1.415L15.485 13H1Zm19.288-7v12h2V6h-2Z"/>`),
		g.Group(children),
	)
}