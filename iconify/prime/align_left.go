package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.93 11h-10a.75.75 0 1 1 0-1.5h10a.75.75 0 0 1 0 1.5Zm6.14-4h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5Zm0 8h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5Zm-6.14 4h-10a.75.75 0 1 1 0-1.5h10a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}