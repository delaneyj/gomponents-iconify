package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4H9.5a3.5 3.5 0 0 0 0 7h.5m4 4V4m-4 11V4M5 19h14M7 21l-2-2l2-2"/>`),
		g.Group(children),
	)
}