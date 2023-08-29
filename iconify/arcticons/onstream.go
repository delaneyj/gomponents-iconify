package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onstream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.909 2.498l11.515 3.223L22.44 22.105l12.016 3.323L4.144 45.18l18.205-17.932l-9.921-2.776L33.909 2.498z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.816 5.672h-.01c-2.35-.94-4.91-1.45-7.59-1.45c-11.4 0-20.64 9.24-20.64 20.64c0 6.04 2.59 11.47 6.72 15.24m29.02-29.31c3.44 3.69 5.54 8.63 5.54 14.07c0 11.4-9.24 20.64-20.64 20.64c-4.82 0-9.25-1.65-12.77-4.43"/>`),
		g.Group(children),
	)
}