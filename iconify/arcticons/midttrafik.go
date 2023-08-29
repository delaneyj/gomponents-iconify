package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Midttrafik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.18 31.98c3.76 2.41 14.75-6.17 14.68-10.56c-.07-4.4-13.75-16.38-20.14-13.05c-6.38 3.33-.29 19.92 5.46 23.61Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.04 38.81c4.12-1.71 3.35-15.63-.3-18.08c-3.65-2.45-21.3 1.95-22.2 9.09c-.91 7.14 16.19 11.61 22.5 8.99Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.47 28.41c-1.43 4.23 9.55 12.82 13.79 11.69c4.25-1.13 12.57-17.3 7.8-22.69c-4.77-5.39-19.4 4.53-21.59 11Z"/>`),
		g.Group(children),
	)
}