package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M13 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm18 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0ZM13 36a4 4 0 1 0 8 0a4 4 0 0 0-8 0Z"/><path stroke-linecap="round" d="M4 36h9m8 0h23M4 12h9m8 0h23"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 4v40"/><path stroke-linecap="round" d="M4 24h27m8 0h5"/></g>`),
		g.Group(children),
	)
}