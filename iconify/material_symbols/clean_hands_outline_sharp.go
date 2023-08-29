package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleanHandsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 11.65l-2-.75q-.05-1.2-.913-2.05T10 8q-.825 0-1.513.413T7.4 9.5H5.25q.425-1.275 1.413-2.2T9 6.1V4H7.5V2H13q.85 0 1.6.275t1.375.75L14.55 4.45q-.35-.2-.738-.325T13 4h-2v2.1q1.725.35 2.862 1.713T15 11v.65ZM14 22.5l-7-1.95V22H1V11h7.95L17 14v2h5v4l-8 2.5ZM3 20h2v-7H3v7Zm10.95.4l5.95-1.85V18h-7.075L9.7 16.95l.6-1.9l2.925.95H15v-.65L8.6 13H7v5.5l6.95 1.9ZM9.225 9.5ZM19 10q-.825 0-1.413-.588T17 8q0-.575.425-1.425T19 4q1.15 1.725 1.575 2.575T21 8q0 .825-.588 1.413T19 10Z"/>`),
		g.Group(children),
	)
}