package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 29l.01-8.23c0-.89.91-1.76 1.84-1.77s1.53.5 1.88 1.77l1.8 6.59c.2.78.38 1.64 1.76 1.64s1.85-.85 1.85-1.64V19m19.03 0h-5.79c-1.03 0-1.99.86-2.01 2.46c-.02 1.6.88 2.6 2.01 2.6h3.51c1.3 0 2.28.8 2.28 2.36S31.8 29 30.4 29h-5.92m-8.89 0l2.78-8.3c.44-1.31.87-1.7 1.67-1.7s1.22.37 1.67 1.7l2.77 8.3m10.14 0l2.78-8.3c.43-1.3.87-1.7 1.67-1.7s1.21.37 1.66 1.7L43.5 29"/>`),
		g.Group(children),
	)
}