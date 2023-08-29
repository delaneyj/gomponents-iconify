package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Money(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMoney0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMoney1" fill="currentColor" fill-rule="nonzero"><path id="feMoney2" d="M4 5h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Zm14 2H6a2 2 0 0 1-2 2v6a2 2 0 0 1 2 2h12a2 2 0 0 1 2-2V9a2 2 0 0 1-2-2ZM8 9h2v6H8V9Zm6 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/></g></g>`),
		g.Group(children),
	)
}