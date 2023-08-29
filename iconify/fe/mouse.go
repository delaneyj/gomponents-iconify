package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMouse0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMouse1" fill="currentColor" fill-rule="nonzero"><path id="feMouse2" d="M12 4a4 4 0 0 0-4 4v8a4 4 0 1 0 8 0V8a4 4 0 0 0-4-4Zm0-2a6 6 0 0 1 6 6v8a6 6 0 1 1-12 0V8a6 6 0 0 1 6-6Zm1 8a1 1 0 0 1-2 0V7a1 1 0 0 1 2 0v3Z"/></g></g>`),
		g.Group(children),
	)
}