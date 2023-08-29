package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRepeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 7h-2.086l1.293-1.293a.999.999 0 1 0-1.414-1.414L10.586 8l3.707 3.707a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L14.414 9H16.5c1.379 0 2.5 1.346 2.5 3s-1.346 3-3 3a1 1 0 1 0 0 2c2.757 0 5-2.243 5-5s-2.019-5-4.5-5zm-8.207 5.293a.999.999 0 0 0 0 1.414L9.586 15H7.5C6.121 15 5 13.654 5 12s1.346-3 3-3a1 1 0 1 0 0-2c-2.757 0-5 2.243-5 5s2.019 5 4.5 5h2.086l-1.293 1.293a.999.999 0 1 0 1.414 1.414L13.414 16l-3.707-3.707a.999.999 0 0 0-1.414 0z"/>`),
		g.Group(children),
	)
}