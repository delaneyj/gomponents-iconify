package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpAngleArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 8c.742 0 1.85-.733 2.78-1.475c1.2-.954 2.247-2.094 3.046-3.401C6.425 2.144 7 .956 7 0m0 0c0 .956.575 2.145 1.174 3.124c.8 1.307 1.847 2.447 3.045 3.401C12.15 7.267 13.26 8 14 8M7 0v11.5c0 6.627 5.373 12 12 12h5"/>`),
		g.Group(children),
	)
}