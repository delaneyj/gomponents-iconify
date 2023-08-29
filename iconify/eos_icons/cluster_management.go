package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClusterManagement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1v10h22V1Zm12 6H4V5h9Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3.69 11.37l1.14-1l-1-1.73l-1.45.49a3.647 3.647 0 0 0-1.08-.63L20 14h-2l-.3 1.49a3.646 3.646 0 0 0-1.08.63l-1.45-.49l-1 1.73l1.14 1a3.337 3.337 0 0 0 0 1.26l-1.14 1l1 1.73l1.45-.49a3.645 3.645 0 0 0 1.08.63L18 24h2l.3-1.49a3.646 3.646 0 0 0 1.08-.63l1.45.49l1-1.73l-1.14-1a3.39 3.39 0 0 0 0-1.27ZM19 21a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Z"/><path fill="currentColor" d="M12 19H4v-2h8.294a7.008 7.008 0 0 1 3.114-4H1v10h12.26A6.962 6.962 0 0 1 12 19Z"/>`),
		g.Group(children),
	)
}