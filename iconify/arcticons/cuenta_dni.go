package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CuentaDni(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M33.491 14.545s-3.772.09-5.069 2.333c-1.193 2.064-.965 3.62-.965 3.62s.027 6.859.08 9.654c.054 2.796-.643 4.023-2.091 4.023h-4.827M5.692 33.1a9.026 9.026 0 0 0 1.05 2.802c.49.845 1.115 1.604 1.787 2.312c.935.984 1.975 1.88 3.152 2.557c1.222.702 2.576 1.157 3.958 1.436a17.03 17.03 0 0 0 4.203.315"/><path d="M36.307 20.217h-5.632s.161-2.816 2.816-2.816s2.816 2.816 2.816 2.816ZM5.5 7.457v7.973m27.026 27.027H40.5a2 2 0 0 0 2-2v-8.293m0-16.573V7.457a2 2 0 0 0-2-2h-8.538m-16.331 0H7.5a2 2 0 0 0-2 2m4.418 10.145s1.77-3.781 6.356-3.298m2.253 5.792h-5.632s.161-2.816 2.816-2.816s2.816 2.816 2.816 2.816Z"/></g>`),
		g.Group(children),
	)
}