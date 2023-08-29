package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Osmand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.66 3.2c-6.1 1.61-12.34 8.13-13.77 14.28A19.64 19.64 0 0 0 19 41.13a3 3 0 0 1 1.4.87l3.24 3.41a.5.5 0 0 0 .7 0h0L27.59 42a3 3 0 0 1 1.41-.86A19.66 19.66 0 0 0 18.66 3.2ZM24 33.88a11.7 11.7 0 1 1 .06 0Z"/>`),
		g.Group(children),
	)
}