package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cromite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="6.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.41 11.419l-7.608 7.607C34.562 20.521 35 22.208 35 24s-.438 3.479-1.198 4.974l7.607 7.607C43.972 33.041 45.5 28.704 45.5 24s-1.528-9.042-4.09-12.581Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.344 32.172A10.946 10.946 0 0 1 24 35c-6.075 0-11-4.925-11-11s4.925-11 11-11a10.94 10.94 0 0 1 7.344 2.828l7.425-7.425C34.916 4.753 29.726 2.5 24 2.5C12.126 2.5 2.5 12.126 2.5 24S12.126 45.5 24 45.5c5.726 0 10.916-2.253 14.769-5.903l-7.425-7.425Z"/>`),
		g.Group(children),
	)
}