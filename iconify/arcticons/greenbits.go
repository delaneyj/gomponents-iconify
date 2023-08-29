package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greenbits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.27 17.86h8.56m-8.56 4h7.42m-7.42 4h5.74"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M30.17 24.73a3.56 3.56 0 0 1-2 6.83l-5.63-1.65l4-13.65l5.63 1.64a3.56 3.56 0 0 1-2 6.83Zm0 0l-5.63-1.64m1.99-6.83l-2-.58m-1.98 14.23l-2-.58m7.72-12.56l.71-2.42m2.08 3.23l.7-2.42m-8.18 17.68l.71-2.42m2.08 3.23l.71-2.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.2 4.86L6.69 11.25V27C6.69 35.44 24 43.5 24 43.5S41.31 35.44 41.31 27V11.25L25.8 4.86a4.68 4.68 0 0 0-3.6 0Z"/>`),
		g.Group(children),
	)
}