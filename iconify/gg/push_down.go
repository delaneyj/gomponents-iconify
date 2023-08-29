package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 1h2v14.485l3.243-3.242l1.414 1.414L12 19.314l-5.657-5.657l1.414-1.414L11 15.485V1Zm7 19.288H6v2h12v-2Z"/>`),
		g.Group(children),
	)
}