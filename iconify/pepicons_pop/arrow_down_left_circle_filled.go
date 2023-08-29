package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopArrowDownLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.646 17.098a1 1 0 0 1-1.087.905l-5.185-.472a1 1 0 1 1 .181-1.991l5.186.471a1 1 0 0 1 .905 1.087Z"/><path d="M8.903 10.354a1 1 0 0 1 1.086.906l.471 5.185a1 1 0 1 1-1.991.181l-.472-5.185a1 1 0 0 1 .906-1.087Z"/><path d="M10.172 15.829a1 1 0 0 1 0-1.415l5.656-5.656a1 1 0 1 1 1.415 1.414l-5.657 5.657a1 1 0 0 1-1.414 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopArrowDownLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}