package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M4 39h40V9H4v30Z"/><path stroke-linecap="round" d="m4 9l20 15L44 9"/><path stroke-linecap="round" d="M24 9H4v15m40 0V9H24"/></g>`),
		g.Group(children),
	)
}