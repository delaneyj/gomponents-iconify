package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalBarOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.425 0-.713-.288T6 20q0-.425.288-.713T7 19h4v-5L3.35 5.4q-.15-.175-.25-.438T3 4.45q0-.6.425-1.025T4.45 3h15.1q.6 0 1.025.425T21 4.475q0 .225-.1.488t-.25.437L13 14v5h4q.425 0 .713.288T18 20q0 .425-.288.713T17 21H7Zm.45-14h9.1l1.8-2H5.65l1.8 2ZM12 12.1L14.775 9h-5.55L12 12.1Zm0 0Z"/>`),
		g.Group(children),
	)
}