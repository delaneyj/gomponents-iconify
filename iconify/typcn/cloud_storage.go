package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 9l-.351.015A5.967 5.967 0 0 0 11 5c-3.309 0-6 2.691-6 6l.001.126A4.007 4.007 0 0 0 2 15c0 2.206 1.794 4 4 4h5v-4.586l-1.293 1.293a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414l2.999-2.999a1 1 0 0 1 1.416 0l2.999 2.999a.999.999 0 1 1-1.414 1.414L13 14.414V19h4c2.757 0 5-2.243 5-5s-2.243-5-5-5z"/>`),
		g.Group(children),
	)
}