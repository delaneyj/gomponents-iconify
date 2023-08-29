package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForAruba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m20.026 22.339l2.14 5.727l2.141-5.727l5.726-2.139l-5.726-2.139l-2.141-5.728l-2.14 5.728L14.3 20.2z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-9.834 4.716l3.669 9.816l9.817 3.668l-9.817 3.668l-3.669 9.816l-3.667-9.816L8.681 20.2l9.817-3.668l3.668-9.816M5.133 39.867h53.734a27.927 27.927 0 0 1-1.483 3.933H6.616a27.758 27.758 0 0 1-1.483-3.933m44.623 13.766H14.244a28.191 28.191 0 0 1-3.921-3.933h43.355a28.253 28.253 0 0 1-3.922 3.933"/>`),
		g.Group(children),
	)
}