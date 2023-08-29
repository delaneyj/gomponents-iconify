package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2zm-3.648 11.711L12.002 19l-3.349-3.289a2.129 2.129 0 0 1 0-3.069a2.224 2.224 0 0 1 3.125 0l.224.219l.224-.219a2.225 2.225 0 0 1 3.126 0a2.129 2.129 0 0 1 0 3.069zM19 9H5V7h14v2z"/>`),
		g.Group(children),
	)
}