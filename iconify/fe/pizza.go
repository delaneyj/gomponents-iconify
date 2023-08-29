package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pizza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePizza0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePizza1" fill="currentColor"><path id="fePizza2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-5-5h2v2l-2-2Zm10 0l-2 2v-2h2ZM9 7v2H7l2-2Zm6 0l2 2h-2V7Zm-2.5 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-5 4a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm9 1a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g></g>`),
		g.Group(children),
	)
}