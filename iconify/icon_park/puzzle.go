package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Puzzle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M4 24V12H13V10C13 6.68629 15.6863 4 19 4C22.3137 4 25 6.68629 25 10V12H34V24H38C41.3137 24 44 26.6863 44 30C44 33.3137 41.3137 36 38 36H34V44H4V36H8C11.3137 36 14 33.3137 14 30C14 26.6863 11.3137 24 8 24H4Z"/>`),
		g.Group(children),
	)
}