package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.25 7A6.25 6.25 0 0 0 4 13.25v21.5A6.25 6.25 0 0 0 10.25 41h27.5A6.25 6.25 0 0 0 44 34.75v-21.5A6.25 6.25 0 0 0 37.75 7h-27.5ZM41.5 34.75a3.75 3.75 0 0 1-3.75 3.75h-19.5v-18H41.5v14.25Zm0-16.75H18.25V9.5h19.5a3.75 3.75 0 0 1 3.75 3.75V18Z"/>`),
		g.Group(children),
	)
}