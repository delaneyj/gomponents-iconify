package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hayago(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="14.25" cy="14.25" r="6.882" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.75" cy="33.75" r="6.882" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.794 24.573h14.912l-4.588 4.015v4.588h1.147v2.295h-1.147v3.44h1.147v2.295h-2.868V43.5h-2.294v-2.294h-2.868v-2.294h1.147V35.47h-1.147v-2.294h1.147v-4.589Zm19.5-1.146h14.912l-4.588-4.015v-4.588h1.147v-2.295h-1.147V9.09h1.147V6.793h-2.868V4.5h-2.294v2.294h-2.868v2.294h1.147v3.441h-1.147v2.295h1.147v4.588Z"/>`),
		g.Group(children),
	)
}