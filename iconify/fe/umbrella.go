package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUmbrella0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUmbrella1" fill="currentColor"><path id="feUmbrella2" d="M13 20a2 2 0 1 1-2-2v-5H3a9 9 0 0 1 8-8.945V3a1 1 0 0 1 2 0v1.055A9 9 0 0 1 21 13h-8v7Zm-7.71-9h13.42a7.003 7.003 0 0 0-13.42 0Z"/></g></g>`),
		g.Group(children),
	)
}