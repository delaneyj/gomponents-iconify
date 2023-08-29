package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleanHandsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 22V11h4v11H1Zm13 0l-7-1.975V11h1.975L17 14v2h-4l-1.75-.675l-.35.925L13 17h9v2l-8 3Zm1-10.35L9.225 9.5H5.25q.425-1.275 1.413-2.2T9 6.1V4H7.5V2H13q.85 0 1.6.275t1.375.75L14.55 4.45q-.35-.2-.738-.325T13 4h-2v2.1q1.725.35 2.862 1.713T15 11v.65ZM19 10q-.825 0-1.413-.588T17 8q0-.575.425-1.425T19 4q1.15 1.725 1.575 2.575T21 8q0 .825-.588 1.413T19 10Z"/>`),
		g.Group(children),
	)
}