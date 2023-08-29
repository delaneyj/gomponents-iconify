package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m21 13l-10 7v24"/><path d="m21 4l10 7v13l7 5v15H21V4Z" clip-rule="evenodd"/><path d="M4 44h40"/></g>`),
		g.Group(children),
	)
}