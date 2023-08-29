package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Purescript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.166 4.6l-1.24 1.24l3.97 3.97l-3.97 3.97l1.24 1.24l4.58-4.6a.87.87 0 0 0 0-1.23zM6.955 6.74l1.87 1.75h8.23l-1.87-1.75zm-2.1 2.24l-4.6 4.6a.87.87 0 0 0 0 1.23l4.6 4.59l1.23-1.24l-3.97-3.97l3.97-3.97l-1.24-1.24zm3.97 2.14l-1.87 1.76h8.23l1.87-1.76zm-1.87 4.39l1.87 1.75h8.23l-1.87-1.75z"/>`),
		g.Group(children),
	)
}