package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleanHands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 22.5l-7-1.95V11h1.95l6.2 2.3q.85.3 1.35 1.05T17 16h-2q-.675 0-1.363-.038T12.3 15.7l-2-.65l-.3.95l1.575.575q.625.25 1.288.338T14.2 17H19q1.125 0 2.063.5T22 19v1l-8 2.5ZM1 22V11h4v11H1ZM9.225 9.5H5.25q.425-1.275 1.413-2.2T9 6.1V4H7.5V2H13q.85 0 1.6.275t1.375.75L14.55 4.45q-.35-.2-.738-.325T13 4h-2v2.1q1.725.35 2.862 1.713T15 11v.65L9.225 9.5ZM19 10q-.825 0-1.413-.588T17 8q0-.575.425-1.425T19 4q1.15 1.725 1.575 2.575T21 8q0 .825-.588 1.413T19 10Z"/>`),
		g.Group(children),
	)
}