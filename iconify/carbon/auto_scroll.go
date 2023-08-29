package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoScroll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M12 16a4 4 0 1 0 4-4a4 4 0 0 0-4 4zm6 0a2 2 0 1 1-2-2a2 2 0 0 1 2 2z" fill="currentColor"/><path d="M16 27.17l-5.6-5.59L9 23l7 7l7-7l-1.41-1.41L16 27.17z" fill="currentColor"/><path d="M16 4.83l5.58 5.57L23 9l-7-7l-7 7l1.41 1.41L16 4.83z" fill="currentColor"/>`),
		g.Group(children),
	)
}