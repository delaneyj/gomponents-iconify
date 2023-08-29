package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AldiNord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.875 10.499L14.5 27.501m10.297-17.002l-6.375 17.002m10.297-17.002l-6.375 17.002m7.117-12.752h.871M27.867 19h4.039m-5.632 4.251H33.5"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.298 38.413V33.04h1.209a2.35 2.35 0 0 1 2.35 2.35v.673a2.35 2.35 0 0 1-2.35 2.35h-1.21Zm-5.305 0V33.04h1.759c.995 0 1.8.808 1.8 1.805s-.806 1.804-1.8 1.804h-1.76m1.76 0l1.759 1.763m-14.369.001V33.04l3.56 5.373V33.04m1.895 3.593a1.78 1.78 0 0 0 3.56 0V34.82a1.78 1.78 0 0 0-3.56 0v1.813Z"/>`),
		g.Group(children),
	)
}