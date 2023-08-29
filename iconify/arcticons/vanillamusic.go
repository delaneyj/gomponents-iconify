package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vanillamusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95ZM9.66 8a1.63 1.63 0 0 1 1.64 1.66h0a1.63 1.63 0 0 1-1.64 1.64h0A1.63 1.63 0 0 1 8 9.66h0A1.63 1.63 0 0 1 9.66 8Zm28.68 0A1.63 1.63 0 0 1 40 9.66h0a1.63 1.63 0 0 1-1.64 1.64h0a1.63 1.63 0 0 1-1.66-1.64h0A1.63 1.63 0 0 1 38.34 8ZM24 9A15 15 0 1 1 9 24A15 15 0 0 1 24 9ZM9.66 36.7a1.64 1.64 0 0 1 1.64 1.64h0A1.63 1.63 0 0 1 9.66 40h0A1.63 1.63 0 0 1 8 38.34h0a1.63 1.63 0 0 1 1.66-1.64Zm28.68 0A1.63 1.63 0 0 1 40 38.34h0A1.63 1.63 0 0 1 38.34 40h0a1.63 1.63 0 0 1-1.64-1.64h0a1.63 1.63 0 0 1 1.64-1.64Z"/><circle cx="24" cy="24" r="5.31" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}