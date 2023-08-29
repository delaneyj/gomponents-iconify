package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path fill-rule="evenodd" d="M11 8a2 2 0 0 1 2 2v7a2 2 0 1 1-4 0v-7a2 2 0 0 1 2-2Z" clip-rule="evenodd"/><path d="M13 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g><path fill-rule="evenodd" d="M10.25 7.75a.75.75 0 0 1 .75.75v9a.75.75 0 0 1-1.5 0v-9a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M11.5 4.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/></g>`),
		g.Group(children),
	)
}