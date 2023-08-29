package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Humerus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.05 22q-1.275 0-2.163-.875T6 18.975q0-.6.238-1.15t.662-.975l2.125-2.075V11q-.025-.35-.3-.575T8.15 10q-1.125-.85-1.638-1.75T6 6.25q0-1.775 1.313-3.013T10.5 2q.425 0 .813.113t.687.262q.325.125.537.25t.438.125q.5 0 1-.35q.275-.15.575-.275T15.225 2q1.2 0 1.988.838T18 4.874q0 .575-.238 1.163T17.05 7.1l-2.025 2.075v5.575l2.1 2.1q.425.425.65.975t.225 1.15q0 1.275-.863 2.15T15 22q-.6 0-1.15-.225t-.975-.65q-.175-.175-.388-.263t-.462-.087q-.25 0-.463.088t-.387.262q-.425.425-.975.65T9.05 22Z"/>`),
		g.Group(children),
	)
}