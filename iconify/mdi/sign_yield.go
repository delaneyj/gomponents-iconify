package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignYield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 3h-17C2.4 3 1.6 4.3 2.2 5.3l8.5 14.9c.3.5.8.8 1.3.8s1-.3 1.3-.8l8.5-14.9c.6-1-.2-2.3-1.3-2.3M12 18.5L4.3 5h15.3L12 18.5m-5.1-12h10.2l-5.1 9l-5.1-9Z"/>`),
		g.Group(children),
	)
}