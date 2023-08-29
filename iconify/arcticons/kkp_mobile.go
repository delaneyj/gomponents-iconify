package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KkpMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.332 25.355c-3.825.943-8.232 2.237-12.459 3.386l.027 5c4.149.01 8.298-.008 12.432-.03v-8.356Zm5.67 8.325c2.52-.013 4.977-.024 7.498-.026h0c-3.273 7.14-5.336 6.283-9.279 6.283H17.372c-11.239 0-9.121.01-12.872-6.283c2.623.049 5.241.079 7.911.078l.055-19.375c.01-3.55 2.5-4.048 5.407-5.941c-.11.788-.27 1.631 0 2.017c7.345-.53 16.168 1.195 20.62-2.402c-1.878 4.074-.775 7.732-10.678 7.732l-9.942.024v7.053l12.243-3.234c4.946-1.306 4.272-1.115 8.326 3.695c-1.123.005-2.171.043-2.44.44v9.939Z"/>`),
		g.Group(children),
	)
}