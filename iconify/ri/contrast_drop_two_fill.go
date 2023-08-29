package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastDropTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.636 6.633L12 .269l6.364 6.364a9 9 0 1 1-12.728 0ZM12 3.097l-4.95 4.95A6.978 6.978 0 0 0 5 12.997h14a6.978 6.978 0 0 0-2.05-4.95L12 3.097Z"/>`),
		g.Group(children),
	)
}