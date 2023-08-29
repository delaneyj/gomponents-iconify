package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brevent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 10.1c2.91-2.58 5.16-4.6 13.33-4.6c7.47 0 16.4 5.57 16.4 16.32s-9.29 14.46-13.81 14.46s-10.85-2.18-10.85-8.65c0-5.34 3.34-7 5.52-8.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.13 29c1.81 2.67 4.28 3.11 6.87 3.11a9 9 0 0 0 8-9.21c0-6.46-5.28-13-12.5-13S6.27 14.06 6.27 25.05A18.64 18.64 0 0 0 9.3 36"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.58 29.09c1.77-.08 5.17-1.05 5.17-5.66a6.94 6.94 0 0 0-7.3-7.29c-5.79 0-12.09 3.26-12.09 13S19.92 42.5 26.14 42.5s11.8-4.6 11.8-4.6"/>`),
		g.Group(children),
	)
}