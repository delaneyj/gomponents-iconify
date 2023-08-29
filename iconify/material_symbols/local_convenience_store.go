package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalConvenienceStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.025 18h3v-1h-2v-1h2v-3h-3v1h2v1h-2v3Zm7 0h1v-5h-1v2h-1v-2h-1v3h2v2Zm6-6.95V19q0 .825-.588 1.413T19.026 21h-14q-.825 0-1.413-.588T3.026 19v-7.95q-.575-.525-.887-1.35t-.013-1.8l1.05-3.4q.2-.65.713-1.075T5.074 3h13.9q.675 0 1.175.413t.725 1.087l1.05 3.4q.3.975-.012 1.775t-.888 1.375Zm-6.8-1.05q.675 0 1.025-.463t.275-1.037l-.55-3.5h-1.95v3.7q0 .525.35.913t.85.387Zm-4.5 0q.575 0 .937-.388t.363-.912V5h-1.95l-.55 3.5q-.1.6.262 1.05t.938.45Zm-4.45 0q.45 0 .787-.325t.413-.825L7.025 5h-1.95l-1 3.35q-.15.5.162 1.075T5.275 10Zm13.5 0q.725 0 1.05-.575t.15-1.075L18.925 5h-1.9l.55 3.85q.075.5.413.825t.787.325Z"/>`),
		g.Group(children),
	)
}