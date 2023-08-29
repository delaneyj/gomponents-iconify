package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-8h12V9h2v4h4v8H3Zm5-3v-2H6v2h2Zm1.5 0h2v-2h-2v2Zm3.5 0h2v-2h-2v2Zm1.25-9.75L12.8 6.8q.65-.6 1.45-.95T16 5.5q.95 0 1.75.35t1.45.95l-1.45 1.45q-.35-.35-.788-.55T16 7.5q-.525 0-.963.2t-.787.55Zm-2.5-2.5l-1.4-1.4q1.1-1.1 2.55-1.725T16 2q1.65 0 3.1.625t2.55 1.725l-1.4 1.4q-.825-.825-1.912-1.288T16 4q-1.25 0-2.337.463T11.75 5.75Z"/>`),
		g.Group(children),
	)
}