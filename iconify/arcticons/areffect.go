package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Areffect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 23.6c1.89-.89 3.77-1.77 4.65-2.13c0 0 8-13.34 10.49-13.31a3.16 3.16 0 0 1 1.16.17s.77-1.55 2.2-1.55s2.8 1.56 2.8 1.56l3.52.21A11.48 11.48 0 0 1 31.6 8a2.63 2.63 0 0 1 2.59 2.59c.11 1.46.17 2.37.17 2.37s5.78 3.64 6 7.36s0 4.17 0 4.17l-1.48 2l-1.19-1.74l-.77 1.86l-1.45-1.8l-1.13 1.68l-.9-1.81L32 26.53l-.41-1.77l-1.41 1.44l-.74-1.64l-1.57.87l-.49-1.61L25.75 25l-.36-2.28l-2 1.19l-.07-2.25l-1.73.8l-.23-1.83a1.64 1.64 0 0 0-2.17.7a1.93 1.93 0 0 0 0 2.48l1.4-.26l.19 2.32l1.64-.42l-.06 2.29h2.09v2.09l1.73-.36l-.09 2.29l1.86-.9l.32 2.28l2-1l.1 2L32.18 33l.54 1.9l1.61-1.77l.52 1.77l1.8-1.71s.8 2.8-.13 3.93s-2.64 2-3.92 1.64s-10.11-5.18-11.14-5.79s-6.53-1.55-9.45.8A56.12 56.12 0 0 0 7.6 37.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.24 10.28a4 4 0 0 1 2.77 2c.31 1 .19 2.41-1 2.32s-2.77-1.44-2.78-2.37s.17-1.95 1.01-1.95Zm-3.65 33.91c1.14-3.44 2.44-7.25 2.72-7.63c.51-.71 3.09-3.1 3.09-3.1m10.37-7.89l.32 2.07l1.42.39l-.77 1.53l1.53.79l-.75 1.31l1.61.24l.2 1.56"/>`),
		g.Group(children),
	)
}