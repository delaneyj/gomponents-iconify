package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpPointingTriangleWithRightHalfBlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M37.038 14H35v43h25.962L62 56L37.038 14Z"/><path fill="#fff" d="M35 14h1v43H11l-1-1l25-42Z"/><path fill="#3F3F3F" d="M37.038 14H35v43h25.962L62 56L37.038 14Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-miterlimit="10" d="M33.545 15.252L11.32 53.748c-1.09 1.89.273 4.252 2.455 4.252h44.452c2.182 0 3.546-2.362 2.455-4.252L38.455 15.252c-1.091-1.89-3.819-1.89-4.91 0Z"/><path stroke-linecap="round" d="M36 14.5V55"/></g>`),
		g.Group(children),
	)
}