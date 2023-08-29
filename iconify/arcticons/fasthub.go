package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fasthub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.14 16.22a2.8 2.8 0 0 0-2.47 3h0a2.8 2.8 0 0 0 2.47 3h0a2.8 2.8 0 0 0 2.47-3a2.8 2.8 0 0 0-2.47-3Zm9.06 0a2.81 2.81 0 0 0-2.47 3a2.8 2.8 0 0 0 2.47 3a2.79 2.79 0 0 0 2.47-3h0a2.8 2.8 0 0 0-2.47-3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.29 19.26a12.61 12.61 0 0 0-25.21 0v2a2.52 2.52 0 0 1-2.52 2.53h0A2.52 2.52 0 0 1 8 21.29a1.92 1.92 0 0 0-1.92-1.92h0a1.92 1.92 0 0 0-1.93 1.92a1 1 0 0 0 0 .17a6.37 6.37 0 0 0 6.36 6.21a6.47 6.47 0 0 0 2.53-.52V37.8a2.1 2.1 0 0 0 4.2 0h0v-3.24h0v-.07a1.4 1.4 0 0 1 2.8 0v.07h0v7.22h0a2.12 2.12 0 0 0 4.16 0h0v-7.22h0v-.07a1.4 1.4 0 0 1 2.8 0v.07h0v7.22h0a2.11 2.11 0 0 0 4.15 0h0v-7.22h0v-.07a1.4 1.4 0 0 1 2.8 0v.07h0v3.24a2 2 0 0 0 2.1 2h0a2 2 0 0 0 2.1-2ZM35.95 4.84l-3.67 3.67M15.41 4.84l3.67 3.67"/>`),
		g.Group(children),
	)
}