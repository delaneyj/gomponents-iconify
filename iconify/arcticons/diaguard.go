package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diaguard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.7 20.969c3.954 3.774 4.214 9.451 2.074 14.383A13.85 13.85 0 0 1 24 43.5a13.85 13.85 0 0 1-12.775-8.148c-2.14-4.932-1.875-10.608 2.08-14.383A72.54 72.54 0 0 0 23.998 4.5a72.67 72.67 0 0 0 10.7 16.469Z"/>`),
		g.Group(children),
	)
}