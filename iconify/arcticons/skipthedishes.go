package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skipthedishes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.314 19.078l-2.638 9.844m2.432 0l2.638-9.844h3.223a2.449 2.449 0 0 1 2.413 3.306a4.62 4.62 0 0 1-4.185 3.306h-3.223m-8.504-6.612l-2.638 9.844m5.291 0L21.389 24l5.363-4.889M21.389 24h-1.238M10.5 27.844a2.073 2.073 0 0 0 2.125 1.078h1.457a3.44 3.44 0 0 0 3.115-2.46h0A1.823 1.823 0 0 0 15.401 24h-1.61a1.823 1.823 0 0 1-1.796-2.461h0a3.44 3.44 0 0 1 3.115-2.461h1.458a2.073 2.073 0 0 1 2.124 1.078"/>`),
		g.Group(children),
	)
}