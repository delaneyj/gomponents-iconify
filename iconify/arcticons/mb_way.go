package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MbWay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 32.813l2.203-17.625l6.61 17.625l6.609-17.625l2.203 17.625M37.094 24a4.406 4.406 0 0 1 0 8.813h-7.27V15.187h7.27a4.406 4.406 0 0 1 0 8.813h0Zm0 0h-3.635m5.69-12.5l-.36-3.322A3 3 0 0 0 35.807 5.5H12.193a3 3 0 0 0-2.982 2.678L8.85 11.5m-2.701 25l-.41 3.785A2 2 0 0 0 7.728 42.5h32.544a2 2 0 0 0 1.988-2.215l-.409-3.785"/>`),
		g.Group(children),
	)
}