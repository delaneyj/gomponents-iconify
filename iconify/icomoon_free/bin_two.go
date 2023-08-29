package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BinTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 16h10l1-11H2zm7-14V0H6v2H1v3l1-1h12l1 1V2h-5zM9 2H7V1h2v1z"/>`),
		g.Group(children),
	)
}