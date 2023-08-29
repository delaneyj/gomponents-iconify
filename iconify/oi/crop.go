package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1 0v1H0v1h1v5h5v1h1V7h1V6H7V1.5l1-1l-.5-.5l-1 1H2V0H1zm1 2h3.5L2 5.5V2zm4 .5V6H2.5L6 2.5z"/>`),
		g.Group(children),
	)
}