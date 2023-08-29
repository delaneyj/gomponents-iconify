package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToDo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.5 16l3.5 3.5l6-6m8 2.5a13 13 0 1 1-26 0a13 13 0 0 1 26 0Z"/>`),
		g.Group(children),
	)
}