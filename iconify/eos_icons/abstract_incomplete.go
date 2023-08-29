package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractIncomplete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.72 11.01l-4.01-7A1.968 1.968 0 0 0 15.98 3H8a1.968 1.968 0 0 0-1.73 1.01L4.74 6.68l-.8 1.39l-1.68 2.94A2.033 2.033 0 0 0 2 12a2.004 2.004 0 0 0 .26.99l1.68 2.94l.8 1.39l1.53 2.67A1.968 1.968 0 0 0 8 21h7.98a1.968 1.968 0 0 0 1.73-1.01l4.01-7a2.004 2.004 0 0 0 .26-.99a1.956 1.956 0 0 0-.26-.99Zm-4.384 5.974l.006-.004l-.004.007ZM20.28 12l-2.935 4.974c-.118.01-.234.026-.355.026A4.994 4.994 0 0 1 12 12.098v-.118A5.004 5.004 0 0 0 7 7a2.96 2.96 0 0 0-.31.03L7.89 5h8.26l4.14 7Z"/>`),
		g.Group(children),
	)
}