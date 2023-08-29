package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTurnedIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3H7c-1.1 0-1.99.9-1.99 2L5 21l7-3l7 3V5c0-1.1-.9-2-2-2zm0 14.97l-4.21-1.81l-.79-.34l-.79.34L7 17.97V5h10v12.97z"/><path fill="currentColor" d="m7 17.97l4.21-1.81l.79-.34l.79.34L17 17.97V5H7z" opacity=".3"/>`),
		g.Group(children),
	)
}