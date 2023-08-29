package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m221.28 215.51l-90.71-31a7.89 7.89 0 0 0-5.38 0l-90.47 31a8 8 0 0 1-9.67-11.44l95.85-168a8 8 0 0 1 14 0l96.09 168a8 8 0 0 1-9.71 11.44Z" opacity=".2"/><path d="M237.9 200.09L141.85 32.18a16 16 0 0 0-27.89 0l-95.89 168a16 16 0 0 0 19.25 22.92l90.47-31h.19l90.68 31a16 16 0 0 0 19.24-23Zm-14 7.84L136 177.86V120a8 8 0 0 0-16 0v57.78l-87.88 30.16l-.12.06l95.86-168L224 208Z"/></g>`),
		g.Group(children),
	)
}