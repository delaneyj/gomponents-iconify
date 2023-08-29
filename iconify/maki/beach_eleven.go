package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeachEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M3.5 1.65v2.82a1.703 1.703 0 0 0-.58-.04c-.28.07-.56.47-.83.57A3.126 3.126 0 0 1 3.5 1.65zm2.31-.62l1.45 2.42a1.216 1.216 0 0 1 .45-.31c.27-.07.7.13 1 .09a3.106 3.106 0 0 0-2.9-2.2zM.984 10h8.99A4.84 4.84 0 0 0 6.9 8.68L5.57 3.74c.46-.04 1.02.06 1.27-.01L5.23 1.04a3.525 3.525 0 0 0-.62.11a2.96 2.96 0 0 0-.61.23v3.1c.25-.08.68-.42 1.08-.6l1.27 4.75q-.39-.03-.84-.03C1.99 8.6.984 10 .984 10z" fill="currentColor"/>`),
		g.Group(children),
	)
}