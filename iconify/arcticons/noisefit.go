package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noisefit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.953 43.011V21.707a6.546 6.546 0 1 1 13.092 0v4.586"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.59 4.752v21.541a6.546 6.546 0 1 1-13.091 0v-4.586"/>`),
		g.Group(children),
	)
}