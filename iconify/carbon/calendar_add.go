package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 22h-6v-6h-2v6h-6v2h6v6h2v-6h6z"/><path fill="currentColor" d="M28 6c0-1.1-.9-2-2-2h-4V2h-2v2h-8V2h-2v2H6c-1.1 0-2 .9-2 2v20c0 1.1.9 2 2 2h8v-2H6V6h4v2h2V6h8v2h2V6h4v8h2V6z"/>`),
		g.Group(children),
	)
}