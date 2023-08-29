package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAddedOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18l-4.2 1.8q-1 .425-1.9-.163T5 17.976V5q0-.825.588-1.413T7 3h6v2H7v12.95l5-2.15l5 2.15V11h2v6.975q0 1.075-.9 1.663t-1.9.162L12 18Zm0-13H7h6h-1Zm5.825 1.175L20.65 3.35q.3-.3.713-.3t.712.3q.3.3.3.712t-.3.713L18.525 8.3q-.3.3-.7.3t-.7-.3L15.7 6.875q-.275-.275-.275-.687t.275-.713q.3-.3.713-.3t.712.3l.7.7Z"/>`),
		g.Group(children),
	)
}