package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniSdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M13.998 18.739L8 21.923V44h32V4H13.998v14.739Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M21 12v6m12-6v6m-6-6v6"/></g>`),
		g.Group(children),
	)
}