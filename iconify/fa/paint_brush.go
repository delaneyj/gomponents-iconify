package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1615 0q70 0 122.5 46.5T1790 163q0 63-45 151q-332 629-465 752q-97 91-218 91q-126 0-216.5-92.5T755 845q0-128 92-212l638-579q59-54 130-54zM706 1034q39 76 106.5 130t150.5 76l1 71q4 213-129.5 347T486 1792q-123 0-218-46.5T115.5 1618T29 1435T0 1215q7 5 41 30t62 44.5t59 36.5t46 17q41 0 55-37q25-66 57.5-112.5t69.5-76t88-47.5t103-25.5t125-10.5z"/>`),
		g.Group(children),
	)
}