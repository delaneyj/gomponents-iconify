package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneAltThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.107 20.294l2.957-2.957a4.903 4.903 0 0 0 1.244-4.87a24.095 24.095 0 0 1-.881-4.384C18.27 6.6 16.977 5.5 15.488 5.5H8.58c-1.777 0-3.145 1.535-2.989 3.304c1.575 17.829 15.777 32.03 33.606 33.606c1.77.156 3.304-1.207 3.304-2.984v-6.16c0-2.248-1.102-3.536-2.583-3.693a24.095 24.095 0 0 1-4.384-.88a4.903 4.903 0 0 0-4.87 1.243l-2.957 2.957"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.393 5.5c0 2.627.325 5.178.938 7.615c3.397 13.5 15.616 23.492 30.169 23.492"/>`),
		g.Group(children),
	)
}