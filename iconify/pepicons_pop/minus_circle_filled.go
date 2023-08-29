package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMinusCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" d="M8 14a1 1 0 1 1 0-2h10a1 1 0 1 1 0 2H8Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMinusCircleFilled0)"/></g>`),
		g.Group(children),
	)
}