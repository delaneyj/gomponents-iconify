package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilkOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 22h26v22H6z"/><path d="M14 38V28l5 6l5-6v10m18-18L30 10M20 6v6l10-2V4L20 6Zm12 16l10-2v21l-10 3V22ZM19.482 12L6 22h26L19.482 12Z"/></g>`),
		g.Group(children),
	)
}