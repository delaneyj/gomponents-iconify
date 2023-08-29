package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreifunkAutoConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.024 22.065l-3.481 3.481l3.59 3.698m-3.59-3.698l9.136.054m-12.121 2.415a8.694 8.694 0 1 1 .004-.036"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.7 23.171a9.994 9.994 0 1 1 0-.043"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.496 23.33a12.121 12.121 0 1 1 .004-.058m-31.648.184l2.303 2.3l-2.371 2.44m2.371-2.44l-6.036.035"/>`),
		g.Group(children),
	)
}