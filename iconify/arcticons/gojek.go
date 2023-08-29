package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gojek(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.898 45.301C9.026 43.188 1.538 33.161 2.6 22.137C3.662 11.113 12.925 2.699 24 2.699s20.338 8.414 21.4 19.438s-6.426 21.051-17.298 23.164"/><circle cx="23.999" cy="24" r="8.6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}