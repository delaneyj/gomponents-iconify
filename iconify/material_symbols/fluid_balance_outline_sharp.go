package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluidBalanceOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v2H4v12h8v2H2Zm2-2V6v12Zm18 5h-5v-3.1q-1.725-.35-2.863-1.712T13 15V8h10v7q0 1.825-1.137 3.188T19 19.9V21h3v2Zm-2.25-9H21v-4h-6v2h.75q.825 0 1.563.375T18.55 13.4q.2.3.525.45t.675.15ZM18 18q.975 0 1.75-.563T20.825 16H19.75q-.825 0-1.563-.363T16.95 14.6q-.225-.275-.537-.437T15.75 14H15v1q0 1.275.863 2.138T18 18Zm-1.05-4.6ZM6 11h5V9H6v2Zm0 4h5v-2H6v2Z"/>`),
		g.Group(children),
	)
}