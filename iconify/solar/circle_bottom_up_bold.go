package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBottomUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.47 21.53a.75.75 0 0 1 0-1.06l7.72-7.72H6.655a.75.75 0 0 1 0-1.5H12a.75.75 0 0 1 .75.75v5.344a.75.75 0 0 1-1.5 0V13.81l-7.72 7.72a.75.75 0 0 1-1.06 0Z" clip-rule="evenodd"/><path d="m3.518 17.3l3.052-3.052a2.25 2.25 0 0 1 .086-4.498H12A2.25 2.25 0 0 1 14.25 12v5.344a2.25 2.25 0 0 1-4.498.086L6.7 20.482A9.953 9.953 0 0 0 12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.947.556 3.763 1.518 5.3Z"/></g>`),
		g.Group(children),
	)
}