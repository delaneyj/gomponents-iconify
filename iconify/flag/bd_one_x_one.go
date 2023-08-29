package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BdOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#006a4e" d="M0 0h512v512H0z"/><circle cx="230" cy="256" r="170.7" fill="#f42a41"/>`),
		g.Group(children),
	)
}