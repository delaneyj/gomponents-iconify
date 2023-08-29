package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareHintFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.5 6a1.5 1.5 0 0 0 0 3h5a1.5 1.5 0 0 0 0-3h-5ZM42 21.5a1.5 1.5 0 0 0-3 0v5a1.5 1.5 0 0 0 3 0v-5Zm-22 19a1.5 1.5 0 0 1 1.5-1.5h5a1.5 1.5 0 0 1 0 3h-5a1.5 1.5 0 0 1-1.5-1.5Zm-11-19a1.5 1.5 0 0 0-3 0v5a1.5 1.5 0 0 0 3 0v-5Zm5-14A1.5 1.5 0 0 0 12.5 6A6.5 6.5 0 0 0 6 12.5a1.5 1.5 0 0 0 3 0A3.5 3.5 0 0 1 12.5 9A1.5 1.5 0 0 0 14 7.5ZM12.5 42a1.5 1.5 0 0 0 0-3A3.5 3.5 0 0 1 9 35.5a1.5 1.5 0 0 0-3 0a6.5 6.5 0 0 0 6.5 6.5ZM34 7.5A1.5 1.5 0 0 1 35.5 6a6.5 6.5 0 0 1 6.5 6.5a1.5 1.5 0 0 1-3 0A3.5 3.5 0 0 0 35.5 9A1.5 1.5 0 0 1 34 7.5ZM35.5 42a1.5 1.5 0 0 1 0-3a3.5 3.5 0 0 0 3.5-3.5a1.5 1.5 0 0 1 3 0a6.5 6.5 0 0 1-6.5 6.5Z"/>`),
		g.Group(children),
	)
}