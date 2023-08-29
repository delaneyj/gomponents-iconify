package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailNew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 24H4L3.997 8.906l11.434 7.916a1 1 0 0 0 1.138 0L28 8.91V18h2V8a2.003 2.003 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v16a2.003 2.003 0 0 0 2 2h15Zm6.799-16L16 14.784L6.201 8Z"/><circle cx="26" cy="24" r="4" fill="currentColor"/>`),
		g.Group(children),
	)
}