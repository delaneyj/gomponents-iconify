package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KissingFaceWithSmilingEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="23" fill="#FCEA2B"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2"><path stroke-linejoin="round" d="M36.415 41.092s10.525 3.196 0 5.564c0 0 10.458 2.988 0 4.913"/><circle cx="36" cy="36" r="23" stroke-linejoin="round"/><path d="M31.694 32.404a4.726 4.726 0 0 0-8.638 0m25.888 0a4.726 4.726 0 0 0-8.638 0"/></g>`),
		g.Group(children),
	)
}