package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrapOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M8 10v28M24 4v12m-8 8h26"/><path stroke-linejoin="round" d="m37.056 19.011l5.037 5.015l-5.037 5.097"/><path d="M24 32v12"/></g>`),
		g.Group(children),
	)
}