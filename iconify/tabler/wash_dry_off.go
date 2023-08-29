package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashDryOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.116 20.127A2.99 2.99 0 0 1 18 21H6a3 3 0 0 1-3-3V6c0-.827.335-1.576.877-2.12M7 3h11a3 3 0 0 1 3 3v11M3 3l18 18"/>`),
		g.Group(children),
	)
}