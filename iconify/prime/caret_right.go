package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18.75a.76.76 0 0 1-.33-.08a.75.75 0 0 1-.42-.67V6a.75.75 0 0 1 .42-.67a.74.74 0 0 1 .78.07l8 6a.75.75 0 0 1 0 1.2l-8 6a.74.74 0 0 1-.45.15ZM8.75 7.5v9l6-4.5Z"/>`),
		g.Group(children),
	)
}