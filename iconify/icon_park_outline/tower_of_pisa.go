package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerOfPisa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 44h40"/><path stroke-linejoin="round" d="m21.25 7.474l15.455 4.142L28 44H11L21.25 7.474Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m19.317 6.957l19.319 5.176M16.21 18.548l19.32 5.176"/><path stroke-linecap="round" d="m25.482 20.932l1.036-3.864"/><path stroke-linecap="round" stroke-linejoin="round" d="m13.105 30.14l19.319 5.175"/><path stroke-linecap="round" d="m22.482 31.932l1.036-3.864"/><rect width="10" height="4" x="25.183" y="4.387" stroke-linecap="round" stroke-linejoin="round" rx="1" transform="rotate(15 25.183 4.387)"/><path stroke-linecap="round" d="m19.482 42.932l1.036-3.864"/></g>`),
		g.Group(children),
	)
}