package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.184 4.912L12 4.75l-7.184.162a2.809 2.809 0 0 0-2.699 2.3a26.484 26.484 0 0 0 0 9.575a2.809 2.809 0 0 0 2.7 2.3L12 19.25l7.184-.162a2.809 2.809 0 0 0 2.699-2.3c.581-3.166.581-6.41 0-9.575a2.809 2.809 0 0 0-2.7-2.3ZM4.85 6.412L12 6.25l7.15.162c.62.014 1.145.461 1.257 1.072c.154.834.264 1.674.332 2.516H3.26a24.89 24.89 0 0 1 .332-2.516A1.309 1.309 0 0 1 4.85 6.412ZM3.181 12c0 1.512.137 3.023.412 4.516c.112.61.637 1.058 1.257 1.072l7.15.162l7.15-.162a1.309 1.309 0 0 0 1.257-1.072c.275-1.493.412-3.004.412-4.516H3.181Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}