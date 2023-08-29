package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorefrontRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.025 21q-.825 0-1.413-.588T3.026 19v-7.95q-.575-.525-.887-1.35t-.013-1.8l1.05-3.4q.2-.65.713-1.075T5.074 3h13.9q.675 0 1.175.413t.725 1.087l1.05 3.4q.3.975-.012 1.775t-.888 1.375V19q0 .825-.587 1.413T19.025 21h-14Zm9.2-11q.675 0 1.025-.463t.275-1.037l-.55-3.5h-1.95v3.7q0 .525.35.913t.85.387Zm-4.5 0q.575 0 .938-.388t.362-.912V5h-1.95l-.55 3.5q-.1.6.262 1.05t.938.45Zm-4.45 0q.45 0 .788-.325t.412-.825L7.025 5h-1.95l-1 3.35q-.15.5.163 1.075T5.275 10Zm13.5 0q.725 0 1.05-.575t.15-1.075L18.925 5h-1.9l.55 3.85q.075.5.413.825t.787.325Z"/>`),
		g.Group(children),
	)
}