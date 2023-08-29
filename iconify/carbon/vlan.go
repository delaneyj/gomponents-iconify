package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vlan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 17v-2H17v-4h2a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2h-6a2.002 2.002 0 0 0-2 2v5a2.002 2.002 0 0 0 2 2h2v4H2v2h6v4H6a2.002 2.002 0 0 0-2 2v5a2.002 2.002 0 0 0 2 2h6a2.002 2.002 0 0 0 2-2v-5a2.002 2.002 0 0 0-2-2h-2v-4h12v4h-2a2.002 2.002 0 0 0-2 2v5a2.002 2.002 0 0 0 2 2h6a2.002 2.002 0 0 0 2-2v-5a2.002 2.002 0 0 0-2-2h-2v-4ZM13 4h6v5h-6Zm-1 24H6v-5h6Zm14 0h-6v-5h6Z"/>`),
		g.Group(children),
	)
}