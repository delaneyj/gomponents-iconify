package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.1 8.9a2.58 2.58 0 1 0-3.65-3.64L22.1 17.61a2.6 2.6 0 0 0-.75 1.82v14h-5.71a.9.9 0 0 0-.88.9a.87.87 0 0 0 .24.67l8.12 8.19a1.12 1.12 0 0 0 1.57 0L32.85 35a.88.88 0 0 0 0-1.25a1 1 0 0 0-.62-.26H26.5v-13Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.91 15.8L13.55 5.41A2.58 2.58 0 0 0 9.9 9.06l11.45 11.35"/>`),
		g.Group(children),
	)
}