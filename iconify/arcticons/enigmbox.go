package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enigmbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.41 40.545h1.18l17.728-5.909L43.5 13.364L24.59 18.09h-1.18L4.5 13.364l1.182 21.272ZM4.5 13.364l18.91-5.91h1.18l18.91 5.91h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.196 30.958a5.545 5.545 0 0 0 7.734-1.29m-7.734-7.734v2.578m7.734-1.288v-2.579m-24.126 9.96a2.082 2.082 0 0 1 .15-1.404a2.054 2.054 0 0 1 .524-.654c.21-.186.442-.348.658-.527a4.867 4.867 0 0 0 1.36-1.785a4.228 4.228 0 0 0 .33-2.213a3.74 3.74 0 0 0-.921-2.03a3.354 3.354 0 0 0-1.95-1.063a3.159 3.159 0 0 0-3.546 2.364"/><circle cx="14.014" cy="32.55" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}