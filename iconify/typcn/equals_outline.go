package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 12H6c-1.654 0-3-1.346-3-3s1.346-3 3-3h12c1.654 0 3 1.346 3 3s-1.346 3-3 3zM6 8a1.001 1.001 0 0 0 0 2h12a1.001 1.001 0 0 0 0-2H6zm12 11H6c-1.654 0-3-1.346-3-3s1.346-3 3-3h12c1.654 0 3 1.346 3 3s-1.346 3-3 3zM6 15a1.001 1.001 0 0 0 0 2h12a1.001 1.001 0 0 0 0-2H6z"/>`),
		g.Group(children),
	)
}