package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M40.5 12.5c0-1.48-.311-2-1.872-2H11.726l-.801-5c-.109-1.46-.85-2-2.421-2H2.501C1.02 3.5.5 3.99.5 5.5v1c0 1.551.52 2 2.001 2h3.722l3.282 19c.35 1.04 1.311 1.95 3.001 2h22.012c1.75 0 2.57-.359 3.002-2l2.98-15zm-7.023 12H13.696l-1.471-9h22.951l-1.699 9zm-19.97 12a4 4 0 0 0 4.002 4a4 4 0 1 0 0-8a4 4 0 0 0-4.002 4zm13.007 0a4 4 0 0 0 4.002 4a4 4 0 1 0 0-8a4 4 0 0 0-4.002 4z"/>`),
		g.Group(children),
	)
}