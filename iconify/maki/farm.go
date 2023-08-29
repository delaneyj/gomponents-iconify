package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Farm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 6L5 4L2 6L1 8v4h2v-2h4v2h2V8zM6 8H4V6h2zm8 4h-3V2.5a1.5 1.5 0 1 1 3 0z"/>`),
		g.Group(children),
	)
}