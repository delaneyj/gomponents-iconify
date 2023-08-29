package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 6V42"/><path d="M7 17.8994L19 5.89941"/><path d="M29 42.1006V6.10059"/><path d="M29 42.1006L41 30.1006"/></g>`),
		g.Group(children),
	)
}