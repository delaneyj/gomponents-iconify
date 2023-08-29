package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonePortraitSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M382 0H130a18 18 0 0 0-18 18v476a18 18 0 0 0 18 18h252a18 18 0 0 0 18-18V18a18 18 0 0 0-18-18ZM148 448V64h216v384Z"/>`),
		g.Group(children),
	)
}