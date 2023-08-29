package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisherOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M16 18a8 8 0 1 1 16 0v24a2 2 0 0 1-2 2H18a2 2 0 0 1-2-2V18Z"/><path stroke-linecap="round" d="M24 24v10"/><path d="M20 5h9v4h-9zm9 0l9-1v6l-9-1V5Z"/><path stroke-linecap="round" d="M20 7c-3 0-7.5-.5-10 2c-2.417 2.416-2 5-2 9"/></g>`),
		g.Group(children),
	)
}