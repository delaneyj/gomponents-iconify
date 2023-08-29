package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4v20.063L8.22 17.28l-1.44 1.44l8.5 8.5l.72.686l.72-.687l8.5-8.5l-1.44-1.44L17 24.063V4h-2z"/>`),
		g.Group(children),
	)
}