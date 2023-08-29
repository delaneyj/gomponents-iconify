package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M2.78 2L2 2.41v12l.78.42l9-6V8l-9-6zM3 13.48V3.35l7.6 5.07L3 13.48z"/><path fill-rule="evenodd" d="m6 14.683l8.78-5.853V8L6 2.147V3.35l7.6 5.07L6 13.48v1.203z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}