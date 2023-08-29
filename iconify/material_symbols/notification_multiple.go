package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V6h2v13h15v2H3Zm4-4q-.825 0-1.413-.588T5 15V6.5q0-.525.238-1t.712-.75L14 0l2.45 1.45l-1.5 1.45l-.95-.55l-6.8 4l6.8 4l6.8-4l-1.35-.8L20.9 4.1l1.15.65q.475.275.712.75t.238 1V15q0 .825-.588 1.413T21 17H7Zm6.85-8l-2.8-2.8l1.4-1.4l1.4 1.4l4.65-4.65l1.4 1.4L13.85 9Z"/>`),
		g.Group(children),
	)
}