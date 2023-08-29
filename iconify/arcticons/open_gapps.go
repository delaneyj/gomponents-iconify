package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenGapps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.654 38.903a2.373 2.373 0 0 0-2.47-2.341a2.458 2.458 0 0 0-2.205 2.428v2.168a2.352 2.352 0 0 0 2.382 2.342h0a2.352 2.352 0 0 0 2.381-2.342h-2.381m11.324 2.34v-6.939h2.293a2.342 2.342 0 1 1 0 4.684h-2.293m-2.156 2.255l-2.294-6.939l-2.381 6.939m.794-2.342h3.087m16.699 1.561a2.118 2.118 0 0 0 1.764.78h1.059a1.755 1.755 0 0 0 1.764-1.734h0a1.755 1.755 0 0 0-1.764-1.734H37.11a1.755 1.755 0 0 1-1.764-1.735h0a1.755 1.755 0 0 1 1.764-1.735h1.058a1.907 1.907 0 0 1 1.765.78m-11.418 6.159v-6.939h2.294a2.342 2.342 0 1 1 0 4.684h-2.294M24 4.5a15.995 15.995 0 0 0-7.998 29.861l4.094-7.086a7.814 7.814 0 1 1 7.824.024L32 34.362A15.996 15.996 0 0 0 24 4.5Z"/>`),
		g.Group(children),
	)
}