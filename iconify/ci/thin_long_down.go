package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 18.5L12 22l3.5-3.5l-.707-.707l-2.293 2.293V2h-1v18.086l-2.293-2.293l-.707.707Z"/>`),
		g.Group(children),
	)
}