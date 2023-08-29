package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaugeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 160v24a16 16 0 0 1-16 16H115.93a4 4 0 0 1-3.24-6.35L174.27 109a8.21 8.21 0 0 0-1.37-11.3a8 8 0 0 0-11.37 1.61l-72 99.06a4 4 0 0 1-3.28 1.63H32a16 16 0 0 1-16-16v-22.87c0-1.79 0-3.57.13-5.33a4 4 0 0 1 4-3.8H48a8 8 0 0 0 8-8.53a8.17 8.17 0 0 0-8.27-7.47H23.92a4 4 0 0 1-3.87-5c12-43.84 49.66-77.13 95.52-82.28a4 4 0 0 1 4.43 4V80a8 8 0 0 0 8.53 8a8.17 8.17 0 0 0 7.47-8.27V52.67a4 4 0 0 1 4.43-4a112.18 112.18 0 0 1 95.8 82.33a4 4 0 0 1-3.88 5h-24.08a8.17 8.17 0 0 0-8.25 7.47a8 8 0 0 0 8 8.53h27.92a4 4 0 0 1 4 3.86c.06 1.37.06 2.75.06 4.14Z"/>`),
		g.Group(children),
	)
}