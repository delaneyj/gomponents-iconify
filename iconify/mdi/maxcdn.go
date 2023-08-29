package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maxcdn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20.6 6.69C19.73 5.61 18.38 5 16.9 5H2.95l1.67 3.57L2.39 19h3.66L8.28 8.57h3.12L9.17 19h3.66l2.23-10.43h1.84c.4 0 .73.13.92.37c.18.23.25.56.18.96L16.04 19h3.65l1.81-8.35c.28-1.44-.04-2.89-.9-3.96z" fill="currentColor"/>`),
		g.Group(children),
	)
}