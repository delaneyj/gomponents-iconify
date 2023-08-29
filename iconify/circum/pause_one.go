package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.25 21.937H6.564a2.5 2.5 0 0 1-2.5-2.5V4.563a2.5 2.5 0 0 1 2.5-2.5H8.25a2.5 2.5 0 0 1 2.5 2.5v14.874a2.5 2.5 0 0 1-2.5 2.5ZM6.564 3.063a1.5 1.5 0 0 0-1.5 1.5v14.874a1.5 1.5 0 0 0 1.5 1.5H8.25a1.5 1.5 0 0 0 1.5-1.5V4.563a1.5 1.5 0 0 0-1.5-1.5Zm10.872 18.874H15.75a2.5 2.5 0 0 1-2.5-2.5V4.563a2.5 2.5 0 0 1 2.5-2.5h1.686a2.5 2.5 0 0 1 2.5 2.5v14.874a2.5 2.5 0 0 1-2.5 2.5ZM15.75 3.063a1.5 1.5 0 0 0-1.5 1.5v14.874a1.5 1.5 0 0 0 1.5 1.5h1.686a1.5 1.5 0 0 0 1.5-1.5V4.563a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}