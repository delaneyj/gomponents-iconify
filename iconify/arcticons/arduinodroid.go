package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arduinodroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.94 20.46a9.15 9.15 0 1 1 2 10c-3.76-4-4.25-9.41-7.72-12.92a9.14 9.14 0 1 0 0 12.92a9.23 9.23 0 0 0 2-2.9M16.07 24h-4.88m25.6 0h-4.88m2.44 2.43v-4.87"/>`),
		g.Group(children),
	)
}