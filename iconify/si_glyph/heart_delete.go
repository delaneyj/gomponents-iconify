package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m11.19 14.2l1.702-1.7l-1.703-1.699l.631-.632l1.701 1.701l1.7-1.7l.631.631l-1.7 1.699l1.7 1.7l-.631.631l-1.7-1.7l-1.699 1.7l-.632-.631z"/><path d="m9.633 10.802l2.188-2.189l1.701 1.702l1.536-1.537c.551-1.094.911-2.326.911-3.716a4.019 4.019 0 0 0-4.011-4.031a4.015 4.015 0 0 0-3.911 3.148a4.054 4.054 0 0 0-3.945-3.148c-2.237 0-4.05 1.824-4.05 4.072c0 6.496 8.005 9.838 8.005 9.838s.729-.304 1.74-.902l1.539-1.538l-1.703-1.699z"/></g>`),
		g.Group(children),
	)
}