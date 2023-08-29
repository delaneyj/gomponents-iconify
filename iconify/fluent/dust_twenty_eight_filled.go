package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DustTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.242 7.736A6.5 6.5 0 1 1 21.19 14a6.5 6.5 0 1 1-10.948 6.264a6.5 6.5 0 1 1 0-12.527ZM8 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm18 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM9 24a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}