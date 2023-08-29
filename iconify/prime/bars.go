package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 12.75H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5Zm0-4.5H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5Zm0 9H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}