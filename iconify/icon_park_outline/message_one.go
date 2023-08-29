package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M4 6h40v30H29l-5 5l-5-5H4V6Z"/><path d="M23 21h2.003m7.998 0H35m-21.999 0H15"/></g>`),
		g.Group(children),
	)
}