package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StyleOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.975 12.85v6.95l-2.7-1.1l2.7-5.85ZM8.9 22H5.975v-8L8.9 22Zm2.325.6L5.375 6.55l11.35-4.15l5.85 16.05l-11.35 4.15Zm-.25-12.6q.425 0 .713-.288T11.975 9q0-.425-.287-.713T10.975 8q-.425 0-.712.288T9.975 9q0 .425.288.713t.712.287Zm1.45 10l7.55-2.75L15.525 5l-7.55 2.75L12.425 20ZM7.975 7.75L15.525 5l-7.55 2.75Z"/>`),
		g.Group(children),
	)
}