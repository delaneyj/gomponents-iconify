package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.022 19.23v12.343m0-5.092a3.085 3.085 0 0 1 3.086-3.085h0a3.085 3.085 0 0 1 3.085 3.085v5.091m-9.522-5.168a3.085 3.085 0 0 1-3.085 3.086h0a3.085 3.085 0 0 1-3.086-3.086V24.4a3.085 3.085 0 0 1 3.086-3.086h0A3.085 3.085 0 0 1 17.67 24.4m.001 5.09v-8.177M36.5 26.404a3.085 3.085 0 0 1-3.086 3.086h0a3.085 3.085 0 0 1-3.085-3.086V24.4a3.085 3.085 0 0 1 3.085-3.086h0A3.085 3.085 0 0 1 36.5 24.4m0 5.09v-8.177M24 19.056l2.634 2.257l3.695-4.886"/>`),
		g.Group(children),
	)
}