package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feColumns0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feColumns1" fill="currentColor" fill-rule="nonzero"><path id="feColumns2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm12 2v12h4V6h-4ZM4 6v12h4V6H4Zm6 0v12h4V6h-4Z"/></g></g>`),
		g.Group(children),
	)
}