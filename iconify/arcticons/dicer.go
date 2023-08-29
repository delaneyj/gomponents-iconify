package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dicer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm5.3 3.65a3.65 3.65 0 1 1-3.65 3.65a3.65 3.65 0 0 1 3.65-3.65Zm22.4 0a3.65 3.65 0 1 1-3.65 3.65a3.66 3.66 0 0 1 3.65-3.65ZM24 20.35A3.64 3.64 0 0 1 27.65 24h0A3.64 3.64 0 0 1 24 27.65h0A3.64 3.64 0 0 1 20.35 24h0A3.64 3.64 0 0 1 24 20.35Zm-11.2 11.2a3.64 3.64 0 0 1 3.65 3.65h0a3.65 3.65 0 0 1-3.65 3.65h0a3.66 3.66 0 0 1-3.65-3.65h0a3.65 3.65 0 0 1 3.65-3.65Zm22.4 0a3.65 3.65 0 0 1 3.65 3.65h0a3.66 3.66 0 0 1-3.65 3.65h0a3.65 3.65 0 0 1-3.65-3.65h0a3.64 3.64 0 0 1 3.65-3.65Z"/>`),
		g.Group(children),
	)
}