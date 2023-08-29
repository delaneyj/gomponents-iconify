package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 21H7a2 2 0 0 1-2-2v-9l7-7l7 7v9a2 2 0 0 1-2 2Z" opacity=".16"/><path d="M20 10a1 1 0 1 0-2 0h2ZM6 10a1 1 0 0 0-2 0h2Zm14.293 2.707a1 1 0 0 0 1.414-1.414l-1.414 1.414ZM12 3l.707-.707a1 1 0 0 0-1.414 0L12 3Zm-9.707 8.293a1 1 0 1 0 1.414 1.414l-1.414-1.414ZM7 22h10v-2H7v2Zm13-3v-9h-2v9h2ZM6 19v-9H4v9h2Zm15.707-7.707l-9-9l-1.414 1.414l9 9l1.414-1.414Zm-10.414-9l-9 9l1.414 1.414l9-9l-1.414-1.414ZM17 22a3 3 0 0 0 3-3h-2a1 1 0 0 1-1 1v2ZM7 20a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Z"/></g>`),
		g.Group(children),
	)
}