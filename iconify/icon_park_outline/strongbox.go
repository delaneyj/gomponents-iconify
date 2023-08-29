package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strongbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><path fill="currentColor" d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 24a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm24-24a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 24a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 33a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/></g>`),
		g.Group(children),
	)
}