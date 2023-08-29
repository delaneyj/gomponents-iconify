package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopExpandCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M8.707 18.707a1 1 0 0 1-1.414-1.414l4-4a1 1 0 1 1 1.414 1.414l-4 4Z"/><path d="M8 19a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2H8Z"/><path d="M9 18a1 1 0 1 1-2 0v-4a1 1 0 1 1 2 0v4Zm5.707-5.293a1 1 0 0 1-1.414-1.414l4-4a1 1 0 1 1 1.414 1.414l-4 4Z"/><path d="M19 12a1 1 0 1 1-2 0V8a1 1 0 1 1 2 0v4Z"/><path d="M14 9a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2h-4Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopExpandCircleFilled0)"/></g>`),
		g.Group(children),
	)
}