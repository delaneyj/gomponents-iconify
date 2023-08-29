package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upscale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.056 1.994l6.91.04l.04 6.91l-2 .011l-.02-3.527l-5.027 5.028l-1.415-1.415l5.027-5.027l-3.526-.02l.011-2ZM2 2h10v2H4v6H2V2Zm0 10h4v2H4v2H2v-4Zm6 0h4v4h-2v-2H8v-2Zm14 0v10h-8v-2h6v-8h2ZM4 18v2h2v2H2v-4h2Zm8 0v4H8v-2h2v-2h2Z"/>`),
		g.Group(children),
	)
}