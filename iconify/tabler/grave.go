package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21v-2a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v2H5zm5-5v-5H6V7h4V3h4v4h4v4h-4v5"/>`),
		g.Group(children),
	)
}