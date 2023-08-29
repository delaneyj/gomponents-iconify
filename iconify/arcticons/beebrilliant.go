package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beebrilliant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 36.44a21.34 21.34 0 0 0 18.5.16m-21.39-7.11c6.8 3.51 15.13 4.72 22.36.81m-4.69-9.79c.13 1 .19 1.76-.82 2.33a8.88 8.88 0 0 1-5 .43c-1.14-.4-1.22-1.06-1.34-2m11.93-2.73c0 8.1 1 14.16-1.35 20.22s-9.77 5.82-14.51 2s-7.64-10.94-7.52-19m2-17l21.67-.29M33.74 20a8.69 8.69 0 0 1-2.86 7.41a6.06 6.06 0 0 1-6.91.74a8.16 8.16 0 0 1-4-6.67l6.88-.65Zm4.82-8.55a3.11 3.11 0 0 1-6.22 0h0a3.1 3.1 0 0 1 3.11-3.1h0a3.1 3.1 0 0 1 3.11 3.08Zm3.75-.11a6.85 6.85 0 1 1-6.84-6.84h0a6.84 6.84 0 0 1 6.84 6.82Zm-19.51 2a8.54 8.54 0 0 1-8.55 8.54h0a8.55 8.55 0 1 1 0-17.09h0a8.55 8.55 0 0 1 8.55 8.55Zm-4.2-.17a4.39 4.39 0 0 1-8.77 0h0a4.38 4.38 0 0 1 4.38-4.38h0a4.38 4.38 0 0 1 4.39 4.35Z"/>`),
		g.Group(children),
	)
}