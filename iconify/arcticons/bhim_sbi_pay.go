package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BhimSbiPay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.878 4.5A2.108 2.108 0 0 0 9.77 6.608v34.784a2.108 2.108 0 0 0 2.108 2.108h24.244a2.108 2.108 0 0 0 2.108-2.108V6.608A2.108 2.108 0 0 0 36.122 4.5ZM9.77 37.8h28.46M24 39.908v-4.216"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.591 26.218l2.4-2.4l-2.4-2.401m2.4 2.401H18.009m2.4-8.725l-2.4 2.4l2.4 2.401m-2.4-2.401h11.982"/>`),
		g.Group(children),
	)
}