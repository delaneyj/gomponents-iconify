package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationMultipleOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V6h2v13h15v2H1Zm13-8.35l-7-4.1V15h14V8.55l-7 4.1ZM5 17V5.3L14 0l2.45 1.45l-1.5 1.45l-.95-.55l-6.8 4l6.8 4l6.8-4l-1.35-.8L20.9 4.1L23 5.275V17H5Zm8.85-8l-2.8-2.8l1.4-1.4l1.4 1.4l4.65-4.65l1.4 1.4L13.85 9Zm.15 6h7H7h7Z"/>`),
		g.Group(children),
	)
}