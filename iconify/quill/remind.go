package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M15 8c-2 2-6.4 5.6-8 4c-3.5-3.5 4.667-3.333 8-4Zm0 0c2.833 0 12-1 8.5 2.5C21 13 17.833 9.667 15 8Zm-2.172 17.828L7.5 20.5a2.121 2.121 0 1 1 3-3l1.83 1.83c.247.247.67.072.67-.277V7a2 2 0 1 1 4 0v7.997c0 .333.419.48.626.22a1.759 1.759 0 0 1 2.747 0l.398.497a.453.453 0 0 0 .557.122l.666-.333a2.249 2.249 0 0 1 2.012 0c.609.304.994.927.994 1.609V23a4 4 0 0 1-4 4h-5.343a4 4 0 0 1-2.829-1.172Z"/>`),
		g.Group(children),
	)
}