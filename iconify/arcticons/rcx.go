package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rcx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.38 5.48c3.27 0 11.51 4.74 8.45 10.47a46.12 46.12 0 0 1-6.15-2.79c-1.13 2 4.51 4.25 4.51 4.25s-.63 2.44-5.89-.27a13.19 13.19 0 0 1-.69 4.55c3.12 1.14 3.69 6.88 1.46 7c-1.82.1-.33-4.39-2.37-4.43c-1.52 3.42-3.7 4.53-4.7 4.92c0 2.19-2.46 4.63-3.95 5.43v2.08c0 3.46 4.18 3 4.18 4.66A1 1 0 0 1 27 42.52h-4.49a2.26 2.26 0 0 1-2.05-2.08v-5.32a7.36 7.36 0 0 1-2.38-3.4a8 8 0 0 1-5.59 2.49c-3.44 0-6.78-1.53-7-6.43a10.25 10.25 0 0 0 6.27 2.68c3.84 0 3.8-2.6 4.09-5.2c.24-2.13 1.1-8.52 8.39-9.21c4.17-.39 3.86-3.35 3.86-5.73s1.63-4.84 5.28-4.84Z"/><circle cx="36.82" cy="10.56" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}