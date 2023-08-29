package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.6 4H8.49L8.2 2.27a1.47 1.47 0 0 0-1.06-1.11A2.799 2.799 0 0 0 6.33 1H4.67a2.798 2.798 0 0 0-.8.14A1.47 1.47 0 0 0 2.8 2.27L2.51 4H1.34a.37.37 0 0 0-.34.49l1.21 4.7c.135.47.56.799 1.05.81h4.48a1.12 1.12 0 0 0 1.06-.81l1.2-4.7A.37.37 0 0 0 9.6 4zM3.27 4l.27-1.61a.68.68 0 0 1 .55-.52c.185-.06.376-.093.57-.1h1.67c.194.007.385.04.57.1a.68.68 0 0 1 .55.52L7.73 4H3.27z" fill="currentColor"/>`),
		g.Group(children),
	)
}