package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewArray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 320V43h64v277H0zM299 43h64v277h-64V43zM85 320V43h192v277H85z"/>`),
		g.Group(children),
	)
}