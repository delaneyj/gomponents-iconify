package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuotaOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.93 12A8.054 8.054 0 1 1 12 5.07V3h-1a10 10 0 1 0 10 10v-1Z"/><path fill="currentColor" d="M20.364 3.636A9.003 9.003 0 0 0 14 1v9h9a9.003 9.003 0 0 0-2.636-6.364ZM16 3.294A7.01 7.01 0 0 1 20.706 8H16Z"/>`),
		g.Group(children),
	)
}