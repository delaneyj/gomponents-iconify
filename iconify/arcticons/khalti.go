package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Khalti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.87 7.45h15.42a1.57 1.57 0 0 1 1.56 1.74l-2.35 22a2.1 2.1 0 0 1-1 1.58l-12.28 7.4a2.6 2.6 0 0 1-2.69 0l-12.29-7.36a2.1 2.1 0 0 1-1-1.58l-2.35-22a1.58 1.58 0 0 1 1.57-1.78h15.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.29 12.88l-5.43.58l1.12 10.45l5.42-.58m10.39-11.17v18.92m10.17 0l-7.79-9.46l7.79-9.4m-7.79 9.4h-2.38"/>`),
		g.Group(children),
	)
}