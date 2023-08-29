package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashtagSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path id="clarityHashtagSolid0" fill="currentColor" d="M31.87 10h-5.55l1-4.83A1 1 0 0 0 26.35 4h-2a1 1 0 0 0-1 .78L22.33 10h-5.4l1-4.83A1 1 0 0 0 17 4h-2a1 1 0 0 0-1 .78L13 10H7a1 1 0 0 0-1 .8l-.41 2a1 1 0 0 0 1 1.2h5.55l-1.64 8h-6a1 1 0 0 0-1 .8l-.41 2a1 1 0 0 0 1 1.2h5.59l-1 4.83a1 1 0 0 0 1 1.17h2a1 1 0 0 0 .95-.78L13.67 26h5.4l-1 4.83A1 1 0 0 0 19 32h2a1 1 0 0 0 1-.78L23.05 26h6a1 1 0 0 0 1-.8l.4-2a1 1 0 0 0-1-1.2h-5.58l1.63-8h6a1 1 0 0 0 1-.8l.41-2a1 1 0 0 0-1.04-1.2Zm-12 12h-5.4l1.64-8h5.4Z"/>`),
		g.Group(children),
	)
}