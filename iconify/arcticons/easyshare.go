package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easyshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.806 24c0 4.186-3.426 7.611-7.612 7.611h-9.083C6.925 31.611 3.5 28.186 3.5 24h0c0-4.186 3.425-7.611 7.611-7.611h9.083c1.384 0 2.685.374 3.807 1.027"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.194 24c0-4.186 3.426-7.611 7.612-7.611h9.083c4.186 0 7.611 3.425 7.611 7.611h0c0 4.186-3.425 7.611-7.611 7.611h-9.083a7.547 7.547 0 0 1-3.807-1.027"/>`),
		g.Group(children),
	)
}