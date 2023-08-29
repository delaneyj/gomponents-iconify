package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabeticalSorting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 4V43.5"/><path d="M7 28H23L7 44H23"/><path d="M7 20L15.2759 4L23 20"/><path d="M44 36L36 44L28 36"/></g>`),
		g.Group(children),
	)
}