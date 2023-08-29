package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SramAxs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.52 14.56l.02 4.28H29.93L20.88 34.5l-2.64-4.46l9.06-15.58l13.22.1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.76 23.26l-.02 4.48h-5.41L26.69 42.5l-2.62-4.25l8.58-14.91l8.12-.08ZM22.11 5.5l2.55 4.33l-8.96 15.64H7.24l.01-4.32l5.8-.03l9.07-15.63Z"/>`),
		g.Group(children),
	)
}