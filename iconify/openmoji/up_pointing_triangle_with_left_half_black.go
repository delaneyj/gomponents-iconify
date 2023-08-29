package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpPointingTriangleWithLeftHalfBlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M34.962 14H37v43H11.038L10 56l24.962-42Z"/><path fill="#fff" d="M37 14h-1v43h25l1-1l-25-42Z"/><path fill="#3F3F3F" d="M34.962 14H37v43H11.038L10 56l24.962-42Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-miterlimit="10" d="M33.545 15.252L11.32 53.748c-1.09 1.89.273 4.252 2.455 4.252h44.452c2.182 0 3.546-2.362 2.455-4.252L38.455 15.252c-1.091-1.89-3.819-1.89-4.91 0Z"/><path stroke-linecap="round" d="M36 14.5V55"/></g>`),
		g.Group(children),
	)
}