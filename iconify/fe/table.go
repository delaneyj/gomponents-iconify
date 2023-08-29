package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTable0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTable1" fill="currentColor" fill-rule="nonzero"><path id="feTable2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm9 10v4h7v-4h-7Zm-9 0v4h7v-4H4Zm9-6v4h7V8h-7ZM4 8v4h7V8H4Z"/></g></g>`),
		g.Group(children),
	)
}