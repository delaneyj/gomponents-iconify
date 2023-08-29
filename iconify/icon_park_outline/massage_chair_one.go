package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageChairOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M15 28V15.652C15 13 18 10 24 10s9 3 9 5.652V28m-21 6v-6h24v6H12Zm8-30h8"/><path d="M8 16v12h32V16M17 43h14m-7-9v9m0-39v6"/></g>`),
		g.Group(children),
	)
}