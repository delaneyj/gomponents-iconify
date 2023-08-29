package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.4 20L6 18.6l5-5V6.875L8.425 9.45L7 8.025l5-5l5.025 5.025L15.6 9.475l-2.6-2.6V14.4L7.4 20Zm9.2.025l-3.2-3.175l1.425-1.425l3.175 3.2l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}