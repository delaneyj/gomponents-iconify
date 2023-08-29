package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 2.5C0 3.3.7 4 1.5 4S3 3.3 3 2.5S2.3 1 1.5 1S0 1.7 0 2.5zm0 5C0 8.3.7 9 1.5 9S3 8.3 3 7.5S2.3 6 1.5 6S0 6.7 0 7.5zm0 5c0 .8.7 1.5 1.5 1.5S3 13.3 3 12.5S2.3 11 1.5 11S0 11.7 0 12.5zM5 1h11v3H5V1zm0 5h11v3H5V6zm0 5h11v3H5v-3z"/>`),
		g.Group(children),
	)
}