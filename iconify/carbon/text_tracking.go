package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTracking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25.29 19.29l-1.41 1.42L26.17 23H5.83l2.29-2.29l-1.41-1.42L2 24l4.71 4.71l1.41-1.42L5.83 25h20.34l-2.29 2.29l1.41 1.42L30 24l-4.71-4.71zM26 17h2L23 3h-2l-5 14h2l1-3h6zm-6.33-5L22 5l2.33 7zM14 3l-4 12L6 3H4l5 14h2l5-14h-2z"/>`),
		g.Group(children),
	)
}