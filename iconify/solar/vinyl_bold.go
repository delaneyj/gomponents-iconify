package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.631 10.488a1.881 1.881 0 1 1 0 3.762a1.881 1.881 0 0 1 0-3.762Z"/><path fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm2.487-18.716a.75.75 0 0 0-.975.716v5.56a3.381 3.381 0 1 0 1.5 2.81V5.093a8.045 8.045 0 0 1 4.273 4.868a.75.75 0 1 0 1.43-.45a9.541 9.541 0 0 0-6.228-6.228Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}