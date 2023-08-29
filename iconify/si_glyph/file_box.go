package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.993.006H3.072l-2.06 8.755v5.17c0 1.334.472 2.028 1.804 2.028h12.28c1.246 0 1.885-.527 1.885-2.111V8.761L14.993.006zM11.016 10H6.985V8.969h4.031V10zM2.505 8.01L3.912.981h10.265L15.54 8.01H2.505z"/><path d="M5 4h7.947v.968H5zm1-2h5.947v.968H6zM4 6h9.951v.965H4z"/></g>`),
		g.Group(children),
	)
}