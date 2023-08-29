package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopControllerCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M15.25 13a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Zm2 2.5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/><path fill-rule="evenodd" d="M17.5 7h-9a5 5 0 0 0-5 5v2a5 5 0 0 0 5 5h9a5 5 0 0 0 5-5v-2a5 5 0 0 0-5-5Zm-12 5a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3v-2Z" clip-rule="evenodd"/><path d="M7 14a1 1 0 1 1 0-2h4a1 1 0 0 1 0 2H7Z"/><path d="M10 15a1 1 0 1 1-2 0v-4a1 1 0 0 1 2 0v4Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopControllerCircleFilled0)"/></g>`),
		g.Group(children),
	)
}