package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audioanchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.65 31c-1.58-.82-4.84-7.19 3.35-7.19c4.9 0 6.61 3.63 6.61 5.6s-2 5-6.61 5S12 30.69 12 25c0-8.46 20.49-7 20.49-16.29c0-2.23-1-4.2-2.62-4.2C24.4 4.5 24 14 24 21.46v22"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.13c4.58 0 9.33.18 13.64-10.29"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.23 34.72a5.87 5.87 0 0 0 4.41-2.88s1 5.46-.62 7.32M24 42.13c-4.58 0-9.33.18-13.64-10.29"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.77 34.72a5.87 5.87 0 0 1-4.41-2.88s-1 5.46.62 7.32m15.17 2.94c-.78 0-2.15.15-2.15 1.4m-2.15-1.4c.78 0 2.15.15 2.15 1.4"/>`),
		g.Group(children),
	)
}