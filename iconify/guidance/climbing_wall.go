package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClimbingWall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m16 4l-3.621 3.621a3 3 0 0 1-2.122.879H6.5l-2.121 2.121a3 3 0 0 0-.879 2.122V16m3-5v4.556a1.945 1.945 0 0 0 3.32 1.374l3.43-3.43h.25l.296.592A7.082 7.082 0 0 0 17.5 17.5m-5.5.498a8.859 8.859 0 0 1-.056-.11l-.194-.388l.25.498Zm0 0A10.623 10.623 0 0 0 17.5 23m2 1V8L17 5.5m-5-5L14.5 3m-4 5.5v5m-2.805-7s-1.81-.557-2.135-1.776a1.768 1.768 0 0 1 1.242-2.163a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61l-.291.079Z"/>`),
		g.Group(children),
	)
}