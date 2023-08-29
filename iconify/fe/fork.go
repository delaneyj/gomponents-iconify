package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFork0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFork1" fill="currentColor"><path id="feFork2" d="M9 7.83V12h3a3 3 0 0 0 3-3V7.83a3.001 3.001 0 1 1 2 0V9a5 5 0 0 1-5 5H9v2.17a3.001 3.001 0 1 1-2 0V7.83a3.001 3.001 0 1 1 2 0ZM8 20a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8-14a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM8 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}