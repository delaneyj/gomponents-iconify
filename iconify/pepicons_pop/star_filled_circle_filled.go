package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFilledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopStarFilledCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" d="m13 18.97l-4.295 1.916a1 1 0 0 1-1.402-1.019l.494-4.677l-3.148-3.493a1 1 0 0 1 .535-1.647l4.6-.976L12.134 5a1 1 0 0 1 1.732 0l2.35 4.074l4.6.976a1 1 0 0 1 .535 1.647l-3.148 3.494l.494 4.676a1 1 0 0 1-1.402 1.018L13 18.971Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopStarFilledCircleFilled0)"/></g>`),
		g.Group(children),
	)
}