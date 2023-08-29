package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDocument0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDocument1" fill="currentColor" fill-rule="nonzero"><path id="feDocument2" d="M15 4H6v16h12V7h-3V4ZM6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm2 9h8v2H8v-2Zm0 4h8v2H8v-2Z"/></g></g>`),
		g.Group(children),
	)
}