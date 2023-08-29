package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="17.779" ry="19.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.079 29.985c7.414-4.334 12.59-2.942 17.059.258"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.396 28.027c-6.612-2.928-12.06-.893-17.258 2.216c-4.69 2.805-9.412 5.297-14.054 5.894"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.458 26.831c2.31-19.391-18.77-20.762-16.96.428"/>`),
		g.Group(children),
	)
}