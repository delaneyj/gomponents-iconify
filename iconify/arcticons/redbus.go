package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redbus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="15.177" cy="32.525" r="4.134"/><circle cx="34.357" cy="31.028" r="4.134"/><path d="m19.076 31.154l11.202-1.213m7.776-.842l4.638-.46c1.964-6.99.273-15.487-4.248-16.68c-5.606-1.48-21.971-.303-30.963 3.532c-4.747 2.096-2.966 9.312-.557 16.979l4.024-.436"/></g>`),
		g.Group(children),
	)
}