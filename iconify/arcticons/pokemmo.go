package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pokemmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.99" cy="23.99" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.83 27.75C4.6 37.83 13.4 45.49 23.99 45.49s19.39-7.66 21.16-17.74h-8.71c-1.61 5.35-6.58 9.24-12.45 9.24s-10.84-3.89-12.45-9.24H2.83Zm-.01-7.5C4.59 10.16 13.4 2.49 23.99 2.49s19.4 7.67 21.17 17.76h-8.71c-1.61-5.36-6.58-9.26-12.46-9.26s-10.85 3.9-12.46 9.26H2.82Z"/>`),
		g.Group(children),
	)
}