package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinLongTwoDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 16l3 3l3-3l-.707-.707l-1.793 1.793V5h-1v12.086l-1.793-1.793L9 16Z"/>`),
		g.Group(children),
	)
}