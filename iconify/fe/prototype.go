package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prototype(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePrototype0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePrototype1" fill="currentColor"><path id="fePrototype2" d="M9 12H6.732a2 2 0 1 1 0-2H9V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-7Zm2 0v7h8V5h-8v5h2.268a2 2 0 1 1 0 2H11Z"/></g></g>`),
		g.Group(children),
	)
}