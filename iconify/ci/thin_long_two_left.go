package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongTwoLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 9l-3 3l3 3l.707-.707L6.914 12.5H19v-1H6.914l1.793-1.793L8 9Z"/>`),
		g.Group(children),
	)
}