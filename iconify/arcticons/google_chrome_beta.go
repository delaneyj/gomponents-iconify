package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleChromeBeta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M39.372 38.5a2 2 0 0 1 0 4h-3.3v-8h3.3a2 2 0 0 1 0 4Zm0 0h-3.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.435 43.324a21.485 21.485 0 1 1 9.89-9.89"/><circle cx="24.179" cy="24.019" r="8.046" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 15.953h19.944M17.21 28.024L7.239 10.752m23.909 17.29l-9.972 17.272"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}