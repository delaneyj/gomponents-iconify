package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bbva(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.191 19.799l-3.348 8l-3.348-8m7.038 6.402l3.348-8l3.348 8m-17.546-2.402a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm0 0h-3.3m-4.313 0a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm.001 0h-3.3"/>`),
		g.Group(children),
	)
}