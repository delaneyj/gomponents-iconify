package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(2)"><ellipse cx="6.43" cy="2.421" rx="6.43" ry="2.421"/><path d="m8.575 11.832l1.014-1.013c-.931.159-1.997.261-3.126.261c-3.516 0-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146a17.9 17.9 0 0 0 2.605-.184l1.176-1.176l-1.669-1.667zM6.494 9.919c3.561 0 6.447-.982 6.447-2.196V4.377c0 .672-2.932 1.674-6.447 1.674S.047 5.049.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196z"/><path d="m14.991 11.799l-.816-.816l-1.702 1.703l-1.668-1.667l-.815.813l1.668 1.668l-1.685 1.686l.814.816l1.687-1.687l1.702 1.703l.815-.814l-1.703-1.702l1.703-1.703z"/></g>`),
		g.Group(children),
	)
}