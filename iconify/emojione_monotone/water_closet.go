package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterCloset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c0 16.566 13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-.465 43h-4.112l-3.821-19.439L19.789 45h-4.204L11 19h3.97l2.895 17.855L21.374 19h4.61l3.364 18.16L32.297 19h3.904l-4.666 26zM46 41.287c1.93 0 3.5-1.67 3.5-3.717H53c0 4.098-3.139 7.43-7 7.43c-3.857 0-7-3.332-7-7.43V26.428C39 22.334 42.143 19 46 19c3.861 0 7 3.334 7 7.428h-3.5c0-2.047-1.57-3.713-3.5-3.713s-3.5 1.666-3.5 3.713V37.57c0 2.047 1.57 3.717 3.5 3.717z"/>`),
		g.Group(children),
	)
}