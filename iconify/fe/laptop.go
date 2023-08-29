package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLaptop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLaptop1" fill="currentColor" fill-rule="nonzero"><path id="feLaptop2" d="M21 20H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2ZM5 6v8h14V6H5Zm0-2h14a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/></g></g>`),
		g.Group(children),
	)
}