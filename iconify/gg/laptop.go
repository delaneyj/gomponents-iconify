package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6Zm2 0h14v8H5V6Z" clip-rule="evenodd"/><path d="M2 18a1 1 0 1 0 0 2h20a1 1 0 1 0 0-2H2Z"/></g>`),
		g.Group(children),
	)
}