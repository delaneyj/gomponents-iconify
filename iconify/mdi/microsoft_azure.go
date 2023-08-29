package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftAzure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.05 4.24L6.56 18.05L2 18l5.09-8.76l5.96-5m.7 1.09L22 19.76H6.74l9.3-1.66l-4.87-5.79l2.58-6.98Z"/>`),
		g.Group(children),
	)
}