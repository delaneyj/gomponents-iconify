package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 17l2.23 4.32L30 22l-3.5 3.167L28 30l-5-2.792L18 30l1.5-4.833L16 22l4.9-.68L23 17z"/><path fill="currentColor" d="M25 5h-3V4a2.006 2.006 0 0 0-2-2h-8a2.006 2.006 0 0 0-2 2v1H7a2.006 2.006 0 0 0-2 2v21a2.006 2.006 0 0 0 2 2h7v-2H7V7h3v3h12V7h3v7h2V7a2.006 2.006 0 0 0-2-2Zm-5 3h-8V4h8Z"/>`),
		g.Group(children),
	)
}