package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15v2h2v-2m10 4v2h2v-2m2-16H5c-1.1 0-2 .9-2 2v4h2V5h14v4h2V5c0-1.1-.9-2-2-2m2 16h-2v2c1.1 0 2-.9 2-2M1 11v2h22v-2M7 19v2h2v-2m10-4v2h2v-2m-10 4v2h2v-2M3 19c0 1.1.9 2 2 2v-2Z"/>`),
		g.Group(children),
	)
}