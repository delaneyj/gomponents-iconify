package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedListBullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1664v-128h128v128H0zm0-384v-128h128v128H0zm0-896h128v128H0V384zm0 512V768h128v128H0z"/>`),
		g.Group(children),
	)
}