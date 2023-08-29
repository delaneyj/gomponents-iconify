package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lesspass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 3.5L44.5 24L24 44.5L3.5 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.998 11.753a5.28 5.28 0 0 0-1.541 10.33v13.67h5.828v-2.742h-2.742v-1.913h2.742v-2.74h-2.742v-6.275a5.281 5.281 0 0 0-1.545-10.33Z"/><circle cx="24" cy="16.998" r="2.42" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}