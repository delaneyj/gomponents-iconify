package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m4.2 8.364l2.05-.228V3.913c0-.643.49-1.18 1.13-1.24a50.256 50.256 0 0 1 9.24 0c.64.06 1.13.597 1.13 1.24v4.223l2.05.228a1.592 1.592 0 0 1 1.369 1.198a14.15 14.15 0 0 1 0 6.82l-.153.615a.662.662 0 0 1-.643.503H17.75v2.583a1.21 1.21 0 0 1-1.09 1.205c-3.099.31-6.221.31-9.32 0a1.21 1.21 0 0 1-1.09-1.205V17.5H3.627a.662.662 0 0 1-.643-.503l-.153-.615a14.15 14.15 0 0 1 0-6.82a1.592 1.592 0 0 1 1.37-1.198Zm3.55-.374a51.753 51.753 0 0 1 8.5 0V4.146a48.756 48.756 0 0 0-8.5 0V7.99Zm0 8.19v3.64a45.2 45.2 0 0 0 8.5 0v-3.64a45.318 45.318 0 0 0-8.5 0Zm10-5.18a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}