package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneBookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 1H8.99C7.89 1 7 1.9 7 3h10c1.1 0 2 .9 2 2v13l2 1V3c0-1.1-.9-2-2-2zm-4 4H5c-1.1 0-2 .9-2 2v16l7-3l7 3V7c0-1.1-.9-2-2-2zm0 14.97l-4.21-1.81l-.79-.34l-.79.34L5 19.97V7h10v12.97z"/><path fill="currentColor" d="m5 19.97l5-2.15l5 2.15V7H5z" opacity=".3"/>`),
		g.Group(children),
	)
}