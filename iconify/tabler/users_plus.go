package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsersPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7a4 4 0 1 0 8 0a4 4 0 0 0-8 0M3 21v-2a4 4 0 0 1 4-4h4c.96 0 1.84.338 2.53.901M16 3.13a4 4 0 0 1 0 7.75M16 19h6m-3-3v6"/>`),
		g.Group(children),
	)
}