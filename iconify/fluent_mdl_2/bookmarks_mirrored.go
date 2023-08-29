package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 512V384h-128v128h128zm-384-128H0v128h1664V384zm384 1280v-128h-128v128h128zm-384 0v-128H0v128h1664zm-512-768V768H0v128h1152zm-512 384v-128H0v128h640z"/>`),
		g.Group(children),
	)
}