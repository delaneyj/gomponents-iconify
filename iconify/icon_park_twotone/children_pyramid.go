package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildrenPyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChildrenPyramid0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M15 17h18v8H15zm-5 8h28v8H10v-8Zm-5 8h38v8H5v-8Z"/><path d="M24 17V7m5 0H19"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChildrenPyramid0)"/>`),
		g.Group(children),
	)
}