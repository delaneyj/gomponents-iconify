package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintbrushSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14.854.146a.5.5 0 0 1 .079.605l-3.841 6.634l-3.477-3.477L14.25.068a.5.5 0 0 1 .605.078ZM6.72 4.427l-1.97 1.14a.5.5 0 0 0-.104.787l4 4a.5.5 0 0 0 .787-.103l1.14-1.97L6.72 4.426ZM.99 10.441a3.063 3.063 0 0 1 2.947-2.227H4a3 3 0 0 1 3 3v.053a2.947 2.947 0 0 1-2.947 2.947h-.08a2.59 2.59 0 0 1-1.115-.252a1.594 1.594 0 0 0-1.57.113l-.51.341a.5.5 0 0 1-.759-.553l.971-3.422Z"/>`),
		g.Group(children),
	)
}