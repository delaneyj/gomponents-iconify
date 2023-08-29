package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.313 0h5.688l-14 24l-14-24h11l3 5.563L17.501 0zM3.5 2L14 20L24.5 2h-3.375L14 14.375L6.875 2z"/>`),
		g.Group(children),
	)
}