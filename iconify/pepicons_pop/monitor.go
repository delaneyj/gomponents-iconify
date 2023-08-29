package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.5 1h-15A2.5 2.5 0 0 0 0 3.5v10A2.5 2.5 0 0 0 2.5 16h15a2.5 2.5 0 0 0 2.5-2.5v-10A2.5 2.5 0 0 0 17.5 1ZM2 3.5a.5.5 0 0 1 .5-.5h15a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-10Z" clip-rule="evenodd"/><path d="M10 13.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5Z"/><path fill-rule="evenodd" d="M11.5 14.5h-3a1 1 0 0 0-1 1V18a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-2.5a1 1 0 0 0-1-1Zm-2.5 3V16h2v1.5H9Z" clip-rule="evenodd"/><path d="M5.5 19.5a1 1 0 1 1 0-2h9a1 1 0 1 1 0 2h-9Z"/><path fill-rule="evenodd" d="M19 12H1v-1h18v1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}