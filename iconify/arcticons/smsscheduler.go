package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smsscheduler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.383 28.274l4.223 4.214a3.622 3.622 0 0 0 5.123.002l.001-.002l8.04-8.04"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.101 22.054a11.026 11.026 0 0 1-10.575 7.916a10.998 10.998 0 0 1-.636-.044v9.277c0 .585.474 1.06 1.06 1.06h24.49a1.06 1.06 0 0 0 1.06-1.06h0v-16.08a1.06 1.06 0 0 0-1.06-1.06H27.1v-.009Z"/><circle cx="16.526" cy="18.945" r="11.026" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.763 13.485a3.448 3.448 0 0 1 5.151-4.586h0m9.55.186a3.445 3.445 0 1 1 4.886 4.788m-9.568-1.528v7.103l3.976 3.384m18.853 15.037l-7.253-5.946m-13.562 5.946l7.254-5.946"/>`),
		g.Group(children),
	)
}