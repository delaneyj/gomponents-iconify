package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 10.75H7a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5Zm3-4H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5Zm0 8H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5Zm-3 4H7a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}