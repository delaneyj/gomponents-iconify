package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleanHandsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 22.5l-7-1.95V22H1V11h7.95l6.2 2.3q.825.3 1.337 1.05T17 16h2q1.25 0 2.125.825T22 19v1l-8 2.5ZM3 20h2v-7H3v7Zm10.95.4l5.95-1.85q-.075-.275-.337-.413T19 18h-4.8q-.775 0-1.4-.1t-1.35-.35l-1.725-.6l.575-1.9l2 .65q.45.15 1.05.225T15 16q0-.275-.137-.513t-.413-.337L8.6 13H7v5.5l6.95 1.9ZM5 16.5Zm10-.5Zm-10 .5Zm2 0Zm8-4.85l-2-.75q-.05-1.2-.913-2.05T10 8q-.825 0-1.513.413T7.4 9.5H5.25q.425-1.275 1.413-2.2T9 6.1V4H7.5V2H13q.85 0 1.6.275t1.375.75L14.55 4.45q-.35-.2-.738-.325T13 4h-2v2.1q1.725.35 2.862 1.713T15 11v.65ZM9.225 9.5ZM19 10q-.825 0-1.413-.588T17 8q0-.575.425-1.425T19 4q1.15 1.725 1.575 2.575T21 8q0 .825-.588 1.413T19 10Z"/>`),
		g.Group(children),
	)
}