package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4a1.999 1.999 0 1 0 0 4a1.999 1.999 0 1 0 0-4zm0 4h-4v2h1.469l-.438 10H10v6H8v2h16v-2h-2v-6h-3.031l-.438-10H20V8zm-.531 2h1.062l.438 10H15.03zM12 22h8v4h-8z"/>`),
		g.Group(children),
	)
}