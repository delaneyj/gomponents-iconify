package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinBigDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 17.5L13 24l6.5-6.5l-.707-.707l-5.293 5.293V0h-1v22.086l-5.293-5.293l-.707.707Z"/>`),
		g.Group(children),
	)
}