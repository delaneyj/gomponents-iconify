package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.788 5.758L9 3h3l4 7h4a2 2 0 1 1 0 4h-2m-2.718 1.256L12 21H9l2-7H7l-2 2H2l2-4l-2-4h3l2 2h3M3 3l18 18"/>`),
		g.Group(children),
	)
}