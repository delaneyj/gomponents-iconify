package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Activity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feActivity0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feActivity1" fill="currentColor"><path id="feActivity2" d="M11 19a1 1 0 0 1-2 0V8a1 1 0 1 1 2 0v11Zm-4 0a1 1 0 0 1-2 0v-7a1 1 0 0 1 2 0v7Zm6 0v-9a1 1 0 0 1 2 0v9a1 1 0 0 1-2 0Zm4-14a1 1 0 0 1 2 0v14a1 1 0 0 1-2 0V5Z"/></g></g>`),
		g.Group(children),
	)
}