package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IgMetall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 5.5l18.134 37H5.866L24 5.5Zm0 4.809v8.968"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.915 24.51c0-1.68-1.458-3.138-3.14-3.026c-1.569.112-2.802 1.57-2.802 3.139v2.802A3.014 3.014 0 0 0 24 30.452h0a3.014 3.014 0 0 0 3.027-3.027H24m-4.484 13.054v-8.968L24 40.479l4.484-8.968v8.968"/>`),
		g.Group(children),
	)
}