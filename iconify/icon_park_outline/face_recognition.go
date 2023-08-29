package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M33 5h8a2 2 0 0 1 2 2v8m0 18v8a2 2 0 0 1-2 2h-8m-18 0H7a2 2 0 0 1-2-2v-8m0-18V7a2 2 0 0 1 2-2h8"/><path d="M24 38c6.627 0 12-6.268 12-14s-5.373-14-12-14s-12 6.268-12 14s5.373 14 12 14Z"/><path stroke-linecap="round" d="M6 24h36m-21.93 6.106c1.267.937 2.577 1.405 3.93 1.405c1.352 0 2.698-.468 4.038-1.405"/></g>`),
		g.Group(children),
	)
}