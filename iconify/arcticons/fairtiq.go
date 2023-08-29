package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fairtiq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="38.915" cy="38.239" opacity=".97" rx="2.869" ry="2.798"/><path d="M26.19 33.098c-4.518 1.793-9.664.26-12.5-3.724a10.83 10.83 0 0 1 .48-13.166a10.542 10.542 0 0 1 12.736-2.785c4.377 2.124 6.772 6.977 5.818 11.789" opacity=".97"/><path d="M32.086 39.562c-8.172 4.869-18.666 3.03-24.704-4.329c-6.039-7.358-5.812-18.032.532-25.127C14.26 3.01 24.821 1.62 32.78 6.83c7.958 5.212 10.928 15.465 6.991 24.137m3.197 3.143c2.422 2.07 2.72 5.672.672 8.104s-5.691 2.802-8.2.834" opacity=".97"/><path d="M26.19 33.111a5.555 5.555 0 0 1 .792-5.878c1.378-1.661 3.598-2.438 5.754-2.013"/><path d="m32.864 25.25l10.153 8.898m-16.884-1.145l9.182 9.886" opacity=".97"/></g>`),
		g.Group(children),
	)
}