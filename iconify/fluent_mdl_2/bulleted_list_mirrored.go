package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedListMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1280v-128h-128v128h128zm0-384V768h-128v128h128zm-384 0V768H0v128h1664zm384-384V384h-128v128h128zm-384-128H0v128h1664V384zm0 896v-128H0v128h1664zm384 384v-128h-128v128h128zm-384 0v-128H0v128h1664z"/>`),
		g.Group(children),
	)
}