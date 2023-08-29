package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedListBulletMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1280v-128h128v128h-128zm0 384v-128h128v128h-128zm0-768V768h128v128h-128zm0-512h128v128h-128V384z"/>`),
		g.Group(children),
	)
}