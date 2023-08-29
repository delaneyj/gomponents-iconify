package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheveronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12.95 10.707l.707-.707L8 4.343L6.586 5.757L10.828 10l-4.242 4.243L8 15.657l4.95-4.95z"/>`),
		g.Group(children),
	)
}