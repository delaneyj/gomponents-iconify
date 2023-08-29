package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.438 4.063H3.563a.5.5 0 1 1 0-1h16.875a.5.5 0 1 1 0 1ZM12.562 8.5h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 8h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1Zm7.874-4H3.562a.5.5 0 1 1 0-1h16.874a.5.5 0 0 1 0 1Zm0 8.437H3.562a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}