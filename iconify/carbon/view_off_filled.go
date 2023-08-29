package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewOffFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.94 15.66a16.4 16.4 0 0 0-5.73-7.45L30 3.41L28.59 2L2 28.59L3.41 30l5.1-5.09A15.38 15.38 0 0 0 16 27a16.69 16.69 0 0 0 14.94-10.66a1 1 0 0 0 0-.68zM16 22.5a6.46 6.46 0 0 1-3.83-1.26L14 19.43A4 4 0 0 0 19.43 14l1.81-1.81A6.49 6.49 0 0 1 16 22.5zm-11.47-.69l5-5A6.84 6.84 0 0 1 9.5 16A6.51 6.51 0 0 1 16 9.5a6.84 6.84 0 0 1 .79.05l3.78-3.77A14.39 14.39 0 0 0 16 5A16.69 16.69 0 0 0 1.06 15.66a1 1 0 0 0 0 .68a15.86 15.86 0 0 0 3.47 5.47z"/>`),
		g.Group(children),
	)
}