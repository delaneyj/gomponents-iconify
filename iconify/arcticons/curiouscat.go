package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Curiouscat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.5c3.61 0 5.24-1.49 7.9-3c2.11-1.19 5.2-1.39 6.47-2.72a12.15 12.15 0 0 0 3.32-9.31c0-1.95-.95-4.83-.95-5.6S44.54 5.5 40 5.5c-3.29 0-10.49 8.17-10.49 8.17A11.34 11.34 0 0 0 24 11.92a11.34 11.34 0 0 0-5.53 1.75S11.27 5.5 8 5.5c-4.52 0-.72 15.62-.72 16.39s-1 3.65-1 5.6a12.15 12.15 0 0 0 3.35 9.31c1.27 1.33 4.36 1.53 6.47 2.72C18.76 41 20.39 42.5 24 42.5Z"/>`),
		g.Group(children),
	)
}