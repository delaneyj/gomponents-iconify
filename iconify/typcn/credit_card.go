package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 7H6c-1.654 0-3 1.346-3 3v7c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3v-7c0-1.654-1.346-3-3-3zm1 10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-4h13v4zm0-6H5v-1a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v1zm-4 5h2a.5.5 0 0 0 0-1h-2a.5.5 0 0 0 0 1z"/>`),
		g.Group(children),
	)
}