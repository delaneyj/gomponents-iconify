package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.37 9.593L8.779 7L1 14.778l2.593 2.593l7.778-7.778ZM15.222 7L23 14.778l-2.576 2.576l-5.202-5.202l-5.202 5.202l-2.576-2.576L15.222 7Z"/>`),
		g.Group(children),
	)
}