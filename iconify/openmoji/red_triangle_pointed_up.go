package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#EA5A47" d="M33.545 15.252L11.32 53.748c-1.09 1.89.273 4.252 2.455 4.252h44.452c2.182 0 3.546-2.362 2.455-4.252L38.455 15.252c-1.091-1.89-3.819-1.89-4.91 0z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" d="M33.545 15.252L11.32 53.748c-1.09 1.89.273 4.252 2.455 4.252h44.452c2.182 0 3.546-2.362 2.455-4.252L38.455 15.252c-1.091-1.89-3.819-1.89-4.91 0z"/>`),
		g.Group(children),
	)
}