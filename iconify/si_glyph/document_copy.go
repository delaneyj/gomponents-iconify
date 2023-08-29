package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentCopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.02.062H5.026V13h10.991V2.991h-5.995V.062zM6.998 10.096V8.937l7.02.096v.997l-7.02.066zm.014-2.12v-.97h7.02v.97h-7.02zm7.005-1.945H6.96V4.938h7.056v1.093z"/><path d="M11.06.062v1.844h4.752L11.06.062zm-8.997 2V16h10.886v-2.021H3.945V2.062H2.063z"/></g>`),
		g.Group(children),
	)
}