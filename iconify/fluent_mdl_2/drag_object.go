package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragObject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1043 1395l90 90l-173 174l-173-174l90-90l83 83l83-83zM877 526l-90-91l173-173l173 173l-90 91l-83-83l-83 83zM269 877l-82 83l82 83l-90 90L6 960l173-173l90 90zm1645 83l-173 173l-90-90l83-83l-83-83l90-90l173 173zM384 640h1152v128H384V640zm0 256h1152v128H384V896zm0 256h1152v128H384v-128z"/>`),
		g.Group(children),
	)
}