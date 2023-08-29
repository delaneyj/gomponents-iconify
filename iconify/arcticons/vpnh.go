package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vpnh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.35 28.12v-8.24h2.7a2.77 2.77 0 0 1 0 5.53h-2.7m7.78 2.71v-8.24l5.46 8.24v-8.24m2.45 0v8.24m5.46-8.24v8.24m-5.46-4.14h5.46m-23.54-4.1l-2.73 8.24l-2.73-8.24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}