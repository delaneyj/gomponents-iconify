package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafFlutteringInWind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b1cc33" d="M28.919 27.91c8.152-6.829 20.773-4.925 20.773-4.925s-.6 12.836-8.624 19.455s-20.774 4.925-20.774 4.925s.473-12.627 8.625-19.455Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29.63 27.91c8.152-6.829 20.773-4.925 20.773-4.925s-.599 12.836-8.624 19.455s-20.774 4.925-20.774 4.925s.473-12.627 8.625-19.455Zm24.456 10.633a10.246 10.246 0 0 1-1.58 4.74a11.317 11.317 0 0 1-7.901 4.741m14.221-6.321a10.245 10.245 0 0 1-1.58 4.74a11.317 11.317 0 0 1-7.901 4.741M17.74 30.642a10.245 10.245 0 0 1 1.58-4.741a11.317 11.317 0 0 1 7.902-4.741M13 27.481a10.245 10.245 0 0 1 1.58-4.74A11.317 11.317 0 0 1 22.481 18M18 50l2.901-2.976"/>`),
		g.Group(children),
	)
}