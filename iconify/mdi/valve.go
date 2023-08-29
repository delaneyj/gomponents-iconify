package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Valve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22H2V2h2m18 0h-2v20h2M17.24 5.34l-4 4a3 3 0 0 0-4 4l-4 4l1.42 1.42l4-4a3 3 0 0 0 4-4l4-4Z"/>`),
		g.Group(children),
	)
}