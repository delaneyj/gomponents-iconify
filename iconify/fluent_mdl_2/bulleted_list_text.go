package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedListText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 1280v-128h1664v128H384zm0 384v-128h1664v128H384zm0-768V768h1664v128H384zm0-512h1664v128H384V384z"/>`),
		g.Group(children),
	)
}