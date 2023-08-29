package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightSightAutoOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19q1.3 0 2.475-.525t2.025-1.5q-3.2-.2-5.35-2.488T8 9q0-.325.025-.638T8.1 7.75q-1.425.8-2.263 2.2T5 13q0 2.5 1.75 4.25T11 19Zm0 2q-3.35 0-5.675-2.325T3 13q0-1.45.475-2.763T4.8 7.888q.85-1.037 2.025-1.724t2.6-.863q.625-.075.975.438t.025 1.062q-.3.5-.363 1.063T10 9q0 2.5 1.75 4.25T16 15q.3 0 .6-.013t.6-.112q.525-.175.913.213t.237.887q-.725 2.35-2.788 3.688T11 21Zm5.4-12l-.5 1.4q-.1.275-.325.438t-.5.162q-.475 0-.738-.388t-.112-.812l2.55-7.175q.1-.275.35-.45t.55-.175h.65q.3 0 .55.175t.35.45l2.55 7.175q.15.425-.112.813t-.738.387q-.275 0-.5-.163T20.1 10.4L19.6 9h-3.2Zm.45-1.35h2.3L18 4l-1.15 3.65Zm-6.675 6.825Z"/>`),
		g.Group(children),
	)
}