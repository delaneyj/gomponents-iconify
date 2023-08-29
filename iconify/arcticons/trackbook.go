package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trackbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.72a2.79 2.79 0 1 0 2.79 2.79A2.79 2.79 0 0 0 24 18.72Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.81 4.5L24.43 8.14a2.08 2.08 0 0 1-.86 0L8.19 4.5a1.86 1.86 0 0 0-1.86 1.86V38a1.86 1.86 0 0 0 1.86 1.86l15.38 3.61a1.79 1.79 0 0 0 .86 0l15.38-3.61A1.86 1.86 0 0 0 41.67 38V6.36a1.86 1.86 0 0 0-1.86-1.86ZM24.46 38a.49.49 0 0 1-.7.06L23.7 38c-1.61-1.86-7.82-9.53-7.82-15.88a8.12 8.12 0 1 1 16.24 0c0 6.27-6.12 13.93-7.66 15.88Z"/>`),
		g.Group(children),
	)
}