package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterOrpOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2q4.025 3.425 6.013 6.363T18 13.8v.2h-2v-.2q0-1.825-1.513-4.125T10 4.65Q7.025 7.375 5.513 9.675T4 13.8q0 .7.125 1.35t.375 1.225v3.475q-1.175-1.125-1.838-2.675T2 13.8q0-2.5 1.988-5.438T10 2Zm0 8.525ZM6.5 22v-6H11v6H6.5ZM8 20.5h1.5v-3H8v3Zm4 1.5v-6h5v3.9h-.9L17 22h-1.5l-.85-2H13.5v2H12Zm6 0v-6h5v4h-3.5v2H18Zm-4.5-3.5h2v-1h-2v1Zm6 0h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}