package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.29a19.5 19.5 0 0 1-39 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".94" d="M39.79 20.29a15.79 15.79 0 0 1-31.58 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M36.07 20.29a12.07 12.07 0 0 1-24.14 0m24.14 0a12.07 12.07 0 1 0-24.14 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M17.41 23.26a2.22 2.22 0 0 0 1.85.83h1.12a1.86 1.86 0 0 0 1.86-1.85h0a1.86 1.86 0 0 0-1.86-1.86h-1.21a1.87 1.87 0 0 1-1.86-1.86h0a1.87 1.87 0 0 1 1.86-1.86h1.12a2 2 0 0 1 1.85.84m4.65-.93V24h3.71"/>`),
		g.Group(children),
	)
}