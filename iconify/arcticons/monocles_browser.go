package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonoclesBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.005 29.023s1.874-1.47 3.021-.909c3.847 1.881 5.53 7.56 13.556 3.215c-5.478 9.967-14.201 5.72-16.577 2.231c-2.375 3.489-11.098 7.736-16.576-2.23c8.026 4.343 9.709-1.335 13.556-3.216c1.147-.56 3.02.909 3.02.909Zm18.735-.66v-6.9"/><circle cx="33.53" cy="17.247" r="5.224" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.261 28.363v-6.9"/><circle cx="14.47" cy="17.247" r="5.224" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="5.262" cy="17.247" r=".75" fill="currentColor"/><circle cx="42.74" cy="17.247" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}