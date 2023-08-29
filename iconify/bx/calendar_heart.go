package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.648 14.711L11.997 18l3.35-3.289a2.129 2.129 0 0 0 0-3.069a2.225 2.225 0 0 0-3.126 0l-.224.219l-.224-.219a2.224 2.224 0 0 0-3.125 0a2.129 2.129 0 0 0 0 3.069z"/><path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2zm.002 16H5V8h14l.002 12z"/>`),
		g.Group(children),
	)
}