package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feShare0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feShare1" fill="currentColor"><path id="feShare2" d="M14.839 14.92a3 3 0 1 1-.8 1.599l-4.873-2.443a3 3 0 1 1 0-4.151l4.873-2.443a3 3 0 1 1 .8 1.599l-4.877 2.438a3.022 3.022 0 0 1 0 .962l4.877 2.438ZM17 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM7 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}