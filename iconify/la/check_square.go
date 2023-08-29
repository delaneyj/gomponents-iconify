package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm14.281 4.281L14 18.562l-3.281-3.28l-1.438 1.437l4 4l.719.687l.719-.687l8-8z"/>`),
		g.Group(children),
	)
}