package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1360 1v1000H0V1h1360zM522 601h318v-77H522v77zm-362-77l317-135v-96L160 159v99l209 84l-209 83v99z"/>`),
		g.Group(children),
	)
}