package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RupeeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.5 8.5V8a.5.5 0 0 0-.325.88L2.5 8.5ZM2 1h11V0H2v1Zm.5 8h3V8h-3v1Zm3-9h-3v1h3V0ZM2.175 8.88l7 6l.65-.76l-7-6l-.65.76ZM10 4.5A4.5 4.5 0 0 0 5.5 0v1A3.5 3.5 0 0 1 9 4.5h1ZM5.5 9A4.5 4.5 0 0 0 10 4.5H9A3.5 3.5 0 0 1 5.5 8v1ZM2 5h11V4H2v1Z"/>`),
		g.Group(children),
	)
}