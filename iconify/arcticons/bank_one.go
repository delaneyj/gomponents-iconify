package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 42.5h37m0-27.935h-37L24 5.5l18.5 9.065zM7 39.447h34M7 17.618h34M26.075 34.953a1.44 1.44 0 0 0 1.44 1.441h0h-7.03h0a1.44 1.44 0 0 0 1.44-1.44V22.11a1.44 1.44 0 0 0-1.44-1.44h7.03a1.44 1.44 0 0 0-1.44 1.44v12.842Zm-12.485 0a1.44 1.44 0 0 0 1.441 1.441h0H8h0a1.44 1.44 0 0 0 1.441-1.44V22.11A1.44 1.44 0 0 0 8 20.67h7.031a1.44 1.44 0 0 0-1.44 1.44v12.842Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.59 34.953a1.44 1.44 0 0 0 1.441 1.441h0H8h0a1.44 1.44 0 0 0 1.441-1.44V22.11A1.44 1.44 0 0 0 8 20.67h7.031a1.44 1.44 0 0 0-1.44 1.44v12.842Zm24.969 0A1.44 1.44 0 0 0 40 36.394h0h-7.031h0a1.44 1.44 0 0 0 1.44-1.44V22.11a1.44 1.44 0 0 0-1.44-1.44H40a1.44 1.44 0 0 0-1.441 1.44v12.842Z"/>`),
		g.Group(children),
	)
}