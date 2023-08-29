package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nintendogamecube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.816 15.126l4.703 2.715v-5.433L6.814 9.695v5.432zm-2.025 1.168l6.73 3.882v3.82l-10.04-5.79V6.616l3.31 1.91v7.769zM12 6.145L7.298 8.863L12 11.579l4.704-2.717L12 6.146zm0-2.332l5.659 3.274l3.31-1.91L12 0L1.975 5.79L5.28 7.695zm7.207 12.48v-3.947l-2.023 1.167v1.614l-4.703 2.715v.005v-5.436L22.518 6.62v11.587L12.48 24v-3.817l6.727-3.887z"/>`),
		g.Group(children),
	)
}