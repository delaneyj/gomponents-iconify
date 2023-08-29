package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithCrossedOutEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FCEA2B" d="M36 13c-12.682 0-23 10.318-23 23s10.318 23 23 23s23-10.318 23-23s-10.318-23-23-23z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="23"/><path stroke-linecap="round" stroke-linejoin="round" d="M29.5 43c1.284-.638 3.985-1.03 6.842-.998c2.624.03 4.99.414 6.158.998m5-14l-5 4m0-4l5 4m-18-4l-5 4m0-4l5 4"/></g>`),
		g.Group(children),
	)
}