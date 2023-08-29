package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m240 112l-36 40H40a8 8 0 0 1-8-8V80a8 8 0 0 1 8-8h164Z" opacity=".2"/><path d="m246 106.65l-36-40a8 8 0 0 0-6-2.65h-68V32a8 8 0 0 0-16 0v32H40a16 16 0 0 0-16 16v64a16 16 0 0 0 16 16h80v64a8 8 0 0 0 16 0v-64h68a8 8 0 0 0 5.95-2.65l36-40a8 8 0 0 0 .05-10.7ZM200.44 144H40V80h160.44l28.8 32Z"/></g>`),
		g.Group(children),
	)
}