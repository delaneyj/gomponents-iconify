package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClusterManagementOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.69 18.37l1.14-1l-1-1.73l-1.45.49a3.647 3.647 0 0 0-1.08-.63L20 14h-2l-.3 1.49a3.646 3.646 0 0 0-1.08.63l-1.45-.49l-1 1.73l1.14 1a3.337 3.337 0 0 0 0 1.26l-1.14 1l1 1.73l1.45-.49a3.645 3.645 0 0 0 1.08.63L18 24h2l.3-1.49a3.646 3.646 0 0 0 1.08-.63l1.45.49l1-1.73l-1.14-1a3.39 3.39 0 0 0 0-1.27ZM19 21a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Zm4-10H1V1h22ZM3 9h18V3H3Zm10-4H4v2h9Zm3 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm3 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M12.294 21H3v-6h10.26a7.026 7.026 0 0 1 2.148-2H1v10h12.26a6.962 6.962 0 0 1-.966-2Z"/><path fill="currentColor" d="M4 19h8a6.994 6.994 0 0 1 .294-2H4Z"/>`),
		g.Group(children),
	)
}