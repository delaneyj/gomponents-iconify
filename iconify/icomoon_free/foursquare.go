package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.306 1.408A1.001 1.001 0 0 0 12.5 1H3a1 1 0 0 0-1 1v12a1.002 1.002 0 0 0 .999 1a1 1 0 0 0 .707-.293L7.413 11h2.586a.999.999 0 0 0 .954-.702l2.5-8a.999.999 0 0 0-.149-.891zM10.515 5H7a1 1 0 0 0 0 2h2.89l-.625 2H7a.997.997 0 0 0-.707.293L4 11.586V3h7.14l-.625 2z"/>`),
		g.Group(children),
	)
}