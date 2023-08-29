package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aidungeon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.545 22.136h2.733m-2.733-9.116h2.733m-1.366 0v9.107m-4.904-3.07H9.136m4.901 3.079l-2.96-9.108l-3.075 9.108M9.48 34.977H8V25.87h1.48a4.555 4.555 0 1 1 0 9.11Zm18.68-3.079h3.075a3.075 3.075 0 1 1-6.15 0v-2.96a3.075 3.075 0 0 1 6.15 0m2.729 6.039V25.87L40 34.977V25.87m-23.455 9.107V25.87l6.036 9.107V25.87"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}