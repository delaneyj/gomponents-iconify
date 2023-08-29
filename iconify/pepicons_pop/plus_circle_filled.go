package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPlusCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M8 14a1 1 0 1 1 0-2h10a1 1 0 1 1 0 2H8Z"/><path d="M12 8a1 1 0 0 1 2 0v10a1 1 0 1 1-2 0V8Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPlusCircleFilled0)"/></g>`),
		g.Group(children),
	)
}