package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.308 12.467a24.095 24.095 0 0 1-.881-4.384C18.27 6.602 16.977 5.5 15.488 5.5H8.58c-1.777 0-3.145 1.535-2.989 3.304c1.575 17.829 15.777 32.03 33.606 33.606c1.77.156 3.304-1.207 3.304-2.984v-6.16c0-2.248-1.102-3.536-2.583-3.693a24.095 24.095 0 0 1-4.384-.88a4.903 4.903 0 0 0-4.87 1.243l-2.957 2.957a31.27 31.27 0 0 1-12.599-12.599l2.957-2.957a4.903 4.903 0 0 0 1.244-4.87Zm4.054 24.77l4.344-4.344m-16.943-8.255l4.344-4.344"/>`),
		g.Group(children),
	)
}