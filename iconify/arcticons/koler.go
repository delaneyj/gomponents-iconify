package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Koler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.72 30.28a62 62 0 0 1-6.3-1.11a3.11 3.11 0 0 0-3 .58c-.53.54-2.06 2.05-3.65 3.61a28.69 28.69 0 0 1-13.1-13.11c1.55-1.59 3-3.11 3.59-3.64a3.11 3.11 0 0 0 .58-3a61.84 61.84 0 0 1-1.11-6.31a2 2 0 0 0-2.3-1.77H6.91A1.5 1.5 0 0 0 5.55 6.9C5 14.59 9.29 22.82 10.16 24.4h0v.06l.12.23h0a35.44 35.44 0 0 0 13 13h0l.44.25h0c2 1.06 9.95 5.06 17.38 4.51a1.5 1.5 0 0 0 1.39-1.36v-8.51a2 2 0 0 0-1.77-2.3Zm-23.01-23L7.24 17.75m9.41 5.93l-4.22 4.22m11.88 3.45l-4.2 4.21m20.61-5.28L30.25 40.75"/>`),
		g.Group(children),
	)
}