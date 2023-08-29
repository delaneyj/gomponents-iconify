package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeepPublicOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.125 9.5ZM8 16H4q-.425 0-.713-.288T3 15v-.575q0-.2.075-.388T3.3 13.7L5 12V5q-.425 0-.713-.288T4 4q0-.425.288-.713T5 3h6.125q-.45.425-.813.925T9.675 5H7v7.85L5.85 14h6.575q.6.35 1.238.588t1.337.337V15q0 .425-.288.713T14 16h-4v5.575q0 .2-.075.388T9.7 22.3l-.35.35q-.15.15-.35.15t-.35-.15l-.35-.35q-.15-.15-.225-.338T8 21.575V16Zm8-8q.625 0 1.063-.438T17.5 6.5q0-.625-.438-1.063T16 5q-.625 0-1.063.438T14.5 6.5q0 .625.438 1.063T16 8Zm0 3q.75 0 1.4-.35t1.075-.975q-.575-.35-1.2-.513T16 9q-.65 0-1.275.163t-1.2.512q.425.625 1.075.975T16 11Zm0 2q-2.075 0-3.538-1.463T11 8q0-2.075 1.463-3.538T16 3q2.075 0 3.538 1.463T21 8q0 2.075-1.463 3.538T16 13Z"/>`),
		g.Group(children),
	)
}