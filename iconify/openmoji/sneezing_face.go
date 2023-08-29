package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SneezingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="39" cy="36" r="24" fill="#FCEA2B"/><path fill="#FFF" d="M8.998 46.982s16.157 1.236 24.952-9.86c.665-.838 1.986-2.267 3.88-.973c0 0 9.323 7.946-10.97 28.667c0 0 .015-4.83-2.377-4.316c-4.388.942-12.096 4.156-17.88 1.628c0 0 14.126-6.337 2.396-15.146z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m28.488 27.413l6.587 4.014l-6.587 3.737m21.664-7.751l-6.587 4.014l6.587 3.737"/><path d="M38.083 58.982c.304.012.61.018.917.018c12.702 0 23-10.298 23-23c0-12.703-10.298-23-23-23c-12.703 0-23 10.297-23 23c0 1.914.234 3.773.674 5.551"/><path d="m37.929 50.003l3.2 1.09l2.95-4.82l3.63 2.85l2.06-2.36"/><path d="M8.998 46.982s16.157 1.236 24.952-9.86c.665-.838 1.986-2.267 3.88-.973c0 0 9.323 7.946-10.97 28.667c0 0 .015-4.83-2.377-4.316c-4.388.942-12.096 4.156-17.88 1.628c0 0 14.126-6.337 2.396-15.146z"/></g>`),
		g.Group(children),
	)
}