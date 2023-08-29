package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BancoEstado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 24.113l13.183-5.788v-4.187h0A17.493 17.493 0 0 0 29.304 8.2h0L42.5 18.975V7.5a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v16.613Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.708 36.337h10.596V8.212h0a17.492 17.492 0 0 1-10.609 5.925h0v4.188L5.5 24.112v4.138l13.195-5.788v-4.137l4.237-1.85v4.1l-4.224 1.875v4.05l4.236-1.85v4.037l-4.236 1.863V26.5L5.5 32.287v4.05l13.208-5.787v5.787Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.708 26.5h0v-4.05l-.013.012h0L5.5 28.25v4.05l13.208-5.8h0zm10.609 9.837H42.5V18.988L29.305 8.213l.012 28.124zm-23.817 0V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-4.162h-37Z"/>`),
		g.Group(children),
	)
}