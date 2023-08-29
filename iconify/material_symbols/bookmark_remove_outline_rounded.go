package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkRemoveOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 7q-.425 0-.713-.288T15 6q0-.425.288-.713T16 5h4q.425 0 .713.288T21 6q0 .425-.288.713T20 7h-4Zm-4 11l-4.2 1.8q-1 .425-1.9-.163T5 17.976V5q0-.825.588-1.413T7 3h6v2H7v12.95l5-2.15l5 2.15V11h2v6.975q0 1.075-.9 1.663t-1.9.162L12 18Zm0-13H7h6h-1Z"/>`),
		g.Group(children),
	)
}