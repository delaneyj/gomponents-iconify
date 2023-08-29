package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFaceWithSmilingEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="23" fill="#fcea2b"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-width="2"><circle cx="36" cy="36" r="23" stroke-linejoin="round"/><path stroke-linejoin="round" d="M45.815 45.227a15.43 15.43 0 0 1-19.63 0"/><path stroke-miterlimit="10" d="M31.694 33.404a4.726 4.726 0 0 0-8.638 0m25.888 0a4.726 4.726 0 0 0-8.638 0"/></g>`),
		g.Group(children),
	)
}