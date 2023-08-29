package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M20 4v40l22-6V10L20 4Z"/><path stroke-linecap="round" d="M6 10h14v28H6V10Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M28 22v4"/></g>`),
		g.Group(children),
	)
}