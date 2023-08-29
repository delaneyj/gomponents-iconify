package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27.34 12.06l-22-8a1 1 0 0 0-1.28 1.28l8 22A1 1 0 0 0 13 28a1 1 0 0 0 .93-.63l3.84-9.6l9.6-3.84a1 1 0 0 0 0-1.87Zm-10.71 4l-.4.16l-.16.4L13 24.2L6.67 6.67L24.2 13Z"/>`),
		g.Group(children),
	)
}