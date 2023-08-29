package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Line(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 21.33c0-8.67-8.73-15.7-19.5-15.7h0C13.74 5.63 5.23 12 4.55 20.28v1C4.5 29.47 12.18 36.16 22 37c2 .1 1.79.31 1.79 1.62h0c0 1.31-.05 2.83-.05 2.83c0 .92.51 1 1.19 1c1.3 0 6.37-3.19 9.07-5.46a30.93 30.93 0 0 0 7.79-9.21h0a13.13 13.13 0 0 0 1.65-5.32c.04-.46.06-.78.06-1.13Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 16.51v8.58h4.29m3.28 0h4.29m-4.29-8.58h4.29m-2.15 0v8.58m13 0h4.29m-4.29-8.58h4.29m-4.29 4.29h2.15m-2.15-4.29v8.58m-8.66 0v-8.58m6.47 8.58v-8.58m-6.47 0l6.47 8.58"/>`),
		g.Group(children),
	)
}