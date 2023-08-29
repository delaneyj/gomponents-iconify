package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 512V384h128v128H0zm384-128h1664v128H384V384zM0 1664v-128h128v128H0zm384 0v-128h1664v128H384zm512-768V768h1152v128H896zm512 384v-128h640v128h-640z"/>`),
		g.Group(children),
	)
}