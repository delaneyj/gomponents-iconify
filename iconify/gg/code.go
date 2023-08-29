package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.953 16.912l-1.36 1.449l-6.562-6.16L8.19 5.64l1.458 1.369l-4.79 5.104l5.094 4.781v.02Zm4.094 0l1.36 1.449l6.562-6.16L15.81 5.64l-1.458 1.369l4.79 5.104l-5.094 4.781v.02Z"/>`),
		g.Group(children),
	)
}