package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brakarbillett(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.837 14.907a33.213 33.213 0 0 1 6.663 2.865a45.104 45.104 0 0 0-4.17 4.313a44.204 44.204 0 0 0-34.8-.358s16.883-.18 17.52 17.683"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.973 39.41s.985-24.1-26.473-26.965c6.4-2.707 15.374-3.956 24.94-2.205c.289-2.71 4.612-1.798 3.727.683a29.305 29.305 0 0 1 5.013 1.432"/>`),
		g.Group(children),
	)
}