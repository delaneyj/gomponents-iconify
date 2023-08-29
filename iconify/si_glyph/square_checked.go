package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m3.017 7.904l1.871-1.873l2.285 2.291l6.092-6.101l1.78 1.783l-7.961 7.973l-4.067-4.073z"/><path d="M12.021 10.116v2.926H1.979V2.958h7.968l1.074-.933H1.002v11.936h11.977L13 9.083l-.979 1.033z"/></g>`),
		g.Group(children),
	)
}