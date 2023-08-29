package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M298.5 21Q369 21 419 71t50 121t-50 121t-120.5 50T178 313t-50-121t50-121t120.5-50zm.5 299q53 0 90.5-37.5T427 192t-37.5-90.5T299 64t-90.5 37.5T171 192t37.5 90.5T299 320zM43 192q0 41 23.5 74t61.5 47v44q-56-14-92-60T0 192T36 87t92-60v44q-38 14-61.5 47T43 192z"/>`),
		g.Group(children),
	)
}