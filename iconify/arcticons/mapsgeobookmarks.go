package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mapsgeobookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a13.101 13.101 0 0 0-13.101 13.101c0 10.253 10.002 22.591 12.623 25.622a.809.809 0 0 0 1.139.08l.08-.08C27.327 40.18 37.1 27.854 37.1 17.6A13.101 13.101 0 0 0 24 4.5Zm0 16.656a4.443 4.443 0 1 1 4.443-4.443v.011A4.443 4.443 0 0 1 24 21.156Z"/>`),
		g.Group(children),
	)
}