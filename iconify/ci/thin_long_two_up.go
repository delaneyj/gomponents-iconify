package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongTwoUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 8l-3-3l-3 3l.707.707L11.5 6.914V19h1V6.914l1.793 1.793L15 8Z"/>`),
		g.Group(children),
	)
}