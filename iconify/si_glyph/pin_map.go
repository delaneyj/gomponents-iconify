package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.953 8.979H6.025V7.668L7 6.991V4.926l-.975-.804V2.021h5.928v2.228l-.974.677v2.065l.975.634l-.001 1.354z"/><path d="M9.986 7.993H8.038v2.614l.989 5.372l.959-5.372V7.993zM6 0h5.959v.942H6z"/></g>`),
		g.Group(children),
	)
}