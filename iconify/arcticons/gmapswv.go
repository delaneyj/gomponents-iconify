package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gmapswv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37.363L37.884 43.5l1.287-1.584L24 4.5L8.829 41.916l1.287 1.584L24 37.363z"/>`),
		g.Group(children),
	)
}