package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M8 12v32h32V4H14l-6 8Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M15 15v6m6-9v6m12-6v6m-6-6v6"/></g>`),
		g.Group(children),
	)
}