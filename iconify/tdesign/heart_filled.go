package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.002 4.818a6.228 6.228 0 0 1 8.51 9.087l-5.225 5.225L12 22.415l-7.28-7.279l-1.23-1.232a6.228 6.228 0 0 1 8.511-9.086Z"/>`),
		g.Group(children),
	)
}