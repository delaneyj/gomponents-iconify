package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm2 30a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm28 2a6 6 0 1 0 0-12a6 6 0 0 0 0 12ZM22 28a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm12-16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/><path d="m11 11l4 4m15-3l-2 2m6 19.5L28 26m-14 5l4-4"/></g>`),
		g.Group(children),
	)
}