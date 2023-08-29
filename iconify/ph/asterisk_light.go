package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.14 179.09a6 6 0 0 1-8.23 2.06L134 138.6V216a6 6 0 0 1-12 0v-77.4l-70.91 42.55A6.07 6.07 0 0 1 48 182a6 6 0 0 1-3.1-11.15L116.34 128L44.91 85.15a6 6 0 0 1 6.18-10.3L122 117.4V40a6 6 0 0 1 12 0v77.4l70.91-42.55a6 6 0 0 1 6.18 10.3L139.66 128l71.43 42.85a6 6 0 0 1 2.05 8.24Z"/>`),
		g.Group(children),
	)
}