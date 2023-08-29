package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shrug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.45 38.56C6.18 35.38 4.5 30.81 4.5 24s1.68-11.38 5.95-14.56m27.1 29.12c4.27-3.18 6-7.75 6-14.56s-1.68-11.38-6-14.56M21.1 12.47a42.1 42.1 0 0 1 3 7m-10.03-5.22a42.81 42.81 0 0 1 3 7.05m16.22-8.11c-1.32 5.54-2 14.92-14.54 20.53"/>`),
		g.Group(children),
	)
}