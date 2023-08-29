package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gpsvc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.555 22.188a3.353 3.353 0 0 0-3.463-3.34A3.484 3.484 0 0 0 8 22.311v3.093a3.34 3.34 0 0 0 6.679 0H11.34m5.414 3.489V18.842h3.267a3.393 3.393 0 0 1 0 6.785h-3.267m8.6 2.135a2.996 2.996 0 0 0 2.513 1.131h1.507a2.513 2.513 0 0 0 0-5.026h-1.633a2.513 2.513 0 0 1 0-5.025h1.508a2.698 2.698 0 0 1 2.513 1.13m2.175 5.531l1.232 3.65l1.187-3.65M40 28.187a1.219 1.219 0 0 1-2.412-.247v-1.22a1.213 1.213 0 0 1 1.218-1.218a1.16 1.16 0 0 1 1.148.959"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}