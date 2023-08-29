package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tibia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9.2L6.875 7.075q-.425-.425-.65-.975T6 4.975q0-1.25.863-2.112T8.95 2h6.1q1.225 0 2.075.863t.85 2.087q0 .625-.25 1.187t-.675.988L15 9.175v5.55l2.1 2.1q.425.425.663.975T18 18.95q0 1.275-.875 2.15t-2.15.875q-.6 0-1.15-.225t-.975-.65q-.175-.2-.387-.288T12 20.726q-.25 0-.463.1t-.387.275q-.425.425-.975.65t-1.15.225q-1.275 0-2.15-.875T6 18.95q0-.6.225-1.15t.65-.975L9 14.75V9.2Z"/>`),
		g.Group(children),
	)
}