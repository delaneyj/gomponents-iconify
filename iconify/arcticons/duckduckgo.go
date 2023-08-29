package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duckduckgo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="17.7" cy="20.51" r="1.63" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27.38" cy="19.51" r="1.46" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2h0a22 22 0 0 0-7.5 42.65a253.51 253.51 0 0 1-5-24.88c0-5.63 1.5-8.66 7.72-9.62a17.81 17.81 0 0 0-6.6-.6c.38-1.27 1.55-1.21 2.28-1.42c.32-.22-1.93-.84-1.85-.8a4.56 4.56 0 0 1 2.89-.65c3.4.15 7.81 2.38 8.22 3.69c6.78.79 8.15 12.13 7.1 13c.69.34 4-1 6.58-1.9C40 20.76 43.1 22.56 39.4 24c0 0-7.15 3.48-10.63 2.74a1.39 1.39 0 0 0-1.95 1.06c.84 2.74 4.28 2.13 6.31 1.89l3.93-.69c6.22-1.31-1.55 6.62-12.71 1.2a17.38 17.38 0 0 0 1.57 7.93c.67-.06 1.22 0 1.28.19a2.49 2.49 0 0 1 .13.49c1.1-.8 4-2.86 4.75-2.69c.92.22 1.1 6.86.31 7.08c-.59.21-3.3-.74-4.73-1.29l.3.65h0c.45.93.89 1.87 1.3 2.81A22 22 0 0 0 24 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.92 38.11a4.17 4.17 0 0 0-2.32.72c-1.76-.84-5.28-2.46-5.32-1.45c-.13 1.32 0 6.81.7 7.21c.53.3 3.52-1.32 5-2.16H24a5.67 5.67 0 0 0 3.29-.43a.41.41 0 0 0 .17-.22l.2.07"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.66 41.89c1.43.55 4.14 1.5 4.73 1.29c.79-.22.61-6.86-.31-7.08c-.75-.17-3.65 1.89-4.75 2.69a2.49 2.49 0 0 0-.13-.49c-.06-.19-.61-.25-1.28-.19"/>`),
		g.Group(children),
	)
}