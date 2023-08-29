package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feedback(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="m4.5 1l-.5.5v1.527a4.551 4.551 0 0 1 1 0V2h9v5h-1.707L11 8.293V7H8.973a4.551 4.551 0 0 1 0 1H10v1.5l.854.354L12.707 8H14.5l.5-.5v-6l-.5-.5h-10Z"/><path fill-rule="evenodd" d="M6.417 10.429a3.5 3.5 0 1 0-3.834 0A4.501 4.501 0 0 0 0 14.5v.5h1v-.5a3.502 3.502 0 0 1 7 0v.5h1v-.5a4.501 4.501 0 0 0-2.583-4.071ZM4.5 10a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}