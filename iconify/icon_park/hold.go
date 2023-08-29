package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 24L26 26C26 26 41 23 43 23C45 23 45 25 43 27C41 29 34 35 28 35C22 35 18 32 14 32C10 32 4 32 4 32"/><path d="M4 20C6 18 10 15 14 15C18 15 27.5 19 29 21C30.5 23 26 26 26 26"/></g>`),
		g.Group(children),
	)
}