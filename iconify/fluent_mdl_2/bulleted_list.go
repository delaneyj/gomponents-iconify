package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1280v-128h128v128H0zm0-384V768h128v128H0zm384 0V768h1664v128H384zM0 512V384h128v128H0zm384-128h1664v128H384V384zm0 896v-128h1664v128H384zM0 1664v-128h128v128H0zm384 0v-128h1664v128H384z"/>`),
		g.Group(children),
	)
}