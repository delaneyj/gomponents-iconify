package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileWord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFileWord0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileWord1" fill="currentColor"><path id="feFileWord2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9 2H6v16h12V7h-3V4Zm-8 8h2l1 3l1-3h2l1 3l1-3h2l-2 6h-2l-1-3l-1 3H9l-2-6Z"/></g></g>`),
		g.Group(children),
	)
}