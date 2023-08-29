package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.437 4.063H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 1 1 0 1ZM16.5 8.5h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 8h-9a.5.5 0 1 1 0-1h9a.5.5 0 1 1 0 1Zm3.937-4H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1Zm0 8.437H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}