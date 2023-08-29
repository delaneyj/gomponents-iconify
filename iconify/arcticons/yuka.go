package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yuka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.022 40.364l15.105-28.122a2.593 2.593 0 0 1 4.02-.893c1.182 1.119 11.295 10.63 12.618 11.873a2.387 2.387 0 0 1-.15 3.834l-28.23 16.13a2.335 2.335 0 0 1-3.363-2.822ZM38.62 4.5l-5.657 10.72l9.156-2.903M17.744 18.54l5.024 4.833m1.904 11.078l-3.994-4.001m-10.059 1.355l2.63 2.502"/>`),
		g.Group(children),
	)
}