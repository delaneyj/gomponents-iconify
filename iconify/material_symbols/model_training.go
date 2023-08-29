package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModelTraining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.15 18.85q-1.025-1.2-1.588-2.688T3 13q0-3.75 2.625-6.375T12 4h.2l-1.6-1.6L12 1l4 4l-4 4l-1.425-1.425L12.15 6H12Q9.1 6 7.05 8.05T5 13q0 1.275.412 2.4t1.163 2.025L5.15 18.85ZM11 18.5q0-.575-.388-1.137t-.862-1.175q-.475-.613-.863-1.275T8.5 13.5q0-1.45 1.025-2.475T12 10q1.45 0 2.475 1.025T15.5 13.5q0 .75-.388 1.413t-.862 1.274q-.475.613-.863 1.175T13 18.5h-2Zm0 2.5v-1.5h2V21h-2Zm7.85-2.15l-1.425-1.425q.75-.9 1.163-2.025T19 13q0-1.65-.687-3.063t-1.888-2.362L17.85 6.15q1.45 1.25 2.3 3.013T21 13q0 1.675-.563 3.163T18.85 18.85Z"/>`),
		g.Group(children),
	)
}