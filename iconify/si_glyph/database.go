package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="6.43" cy="2.421" rx="6.43" ry="2.421"/><path d="M6.463 11.08c-3.516 0-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146s6.447-.959 6.447-2.146V9.444c0 .656-2.931 1.636-6.447 1.636z"/><path d="M6.494 6.051C2.978 6.051.047 5.049.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196s6.447-.982 6.447-2.196V4.377c0 .672-2.931 1.674-6.447 1.674z"/></g>`),
		g.Group(children),
	)
}