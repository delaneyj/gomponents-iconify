package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19 8.974v6.012C19 18.86 15.866 22 12 22s-7-3.14-7-7.014V8.974A7.01 7.01 0 0 1 11.25 2v3.385c-.591.282-1 .886-1 1.585v2.004c0 .969.784 1.754 1.75 1.754s1.75-.785 1.75-1.754V6.97c0-.7-.409-1.303-1-1.585V2A7.01 7.01 0 0 1 19 8.974Z" opacity=".5"/><path d="M13.75 8.974V6.97c0-.7-.409-1.303-1-1.585V2a7.065 7.065 0 0 0-1.5 0v3.385c-.591.282-1 .885-1 1.585v2.004c0 .968.784 1.753 1.75 1.753s1.75-.785 1.75-1.753Z"/></g>`),
		g.Group(children),
	)
}