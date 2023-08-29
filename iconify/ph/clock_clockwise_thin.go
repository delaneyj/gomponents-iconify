package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockClockwiseThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 80v45.74l38.06 22.83a4 4 0 0 1-4.12 6.86l-40-24A4 4 0 0 1 124 128V80a4 4 0 0 1 8 0Zm92-20a4 4 0 0 0-4 4v28.85C211.33 82.46 203 73 193.05 63a92 92 0 1 0-1.9 132a4 4 0 0 0-5.5-5.82a84 84 0 1 1 1.73-120.5C197.7 79 206.39 89 215.53 100H184a4 4 0 0 0 0 8h40a4 4 0 0 0 4-4V64a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}