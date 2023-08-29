package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 10.75H10a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5Zm.07-4h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5Zm0 8h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5Zm-.07 4H10a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}