package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Benz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6.68 34L24 24m0-20v20V4Zm17.32 30L24 24l17.32 10Z"/><path d="M18 4.916A19.992 19.992 0 0 1 24 4c2.09 0 4.106.32 6 .916M4.63 29a19.876 19.876 0 0 0 2.045 5a20.077 20.077 0 0 0 3.042 4m33.653-9a19.88 19.88 0 0 1-2.046 5a20.08 20.08 0 0 1-3.041 4"/></g>`),
		g.Group(children),
	)
}