package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabeticalSortingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 4V43.5"/><path d="M7 4H23L7 20H23"/><path d="M7 44L15.2759 28L23 44"/><path d="M44 36L36 44L28 36"/></g>`),
		g.Group(children),
	)
}