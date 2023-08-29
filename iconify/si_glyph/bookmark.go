package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.677 0H3.326c-.732 0-1.324.605-1.324 1.353v13.294c0 .748.593 1.353 1.324 1.353l4.676-4.013L12.676 16c.732 0 1.326-.605 1.326-1.353V1.353A1.336 1.336 0 0 0 12.677 0zm-2.195 10L8.01 8.633L5.538 10l.473-2.896L4.01 5.055l2.764-.422L8.01 2l1.236 2.633l2.764.422l-2 2.05l.472 2.895z"/>`),
		g.Group(children),
	)
}