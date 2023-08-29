package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.505 17.01L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.677 10.184M20 21l2-2l-2-2m-3 0l-2 2l2 2"/>`),
		g.Group(children),
	)
}