package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplaySettingsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 15H9v-4H7.5v1.25H6v1.5h1.5V15Zm2.5-1.25h8v-1.5h-8v1.5ZM15 11h1.5V9.75H18v-1.5h-1.5V7H15v4ZM6 9.75h8v-1.5H6v1.5ZM8 21v-2H2V3h20v16h-6v2H8Zm-4-4h16V5H4v12Zm0 0V5v12Z"/>`),
		g.Group(children),
	)
}