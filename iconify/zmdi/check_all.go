package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M375 85L240 221l-30-30L345 55zm91-30l30 30l-256 256l-119-119l30-30l89 89zM0 222l30-30l119 119l-30 30z"/>`),
		g.Group(children),
	)
}