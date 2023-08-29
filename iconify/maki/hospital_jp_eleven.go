package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalJpEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.496 1.75H1.504A.505.505 0 0 0 1 2.257v4.53a.507.507 0 0 0 .3.463l3.995 1.956a.501.501 0 0 0 .41 0L9.7 7.213A.507.507 0 0 0 10 6.75V2.257a.505.505 0 0 0-.504-.507zM7.75 6h-1.5v1.5h-1.5V6h-1.5V4.5h1.5V3h1.5v1.5h1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}