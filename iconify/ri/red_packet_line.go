package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedPacketLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.178 9.766a9.981 9.981 0 0 0 4.827-2.622V4.003h-14v3.141a9.98 9.98 0 0 0 4.827 2.622a2.5 2.5 0 0 1 4.346 0Zm.208 2a2.501 2.501 0 0 1-4.762 0a11.94 11.94 0 0 1-4.62-2.015v10.252h14V9.75a11.94 11.94 0 0 1-4.618 2.016ZM4.005 2.004h16a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1v-18a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}