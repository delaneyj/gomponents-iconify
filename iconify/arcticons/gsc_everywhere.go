package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GscEverywhere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.569 18.317V28.92a3.534 3.534 0 0 1-3.535 3.534h0a3.523 3.523 0 0 1-2.499-1.035"/><rect width="7.069" height="9.366" x="11.5" y="18.317" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.534" transform="rotate(-180 15.034 23)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.326 26.893a3.975 3.975 0 0 0 2.907.79h.793a2.339 2.339 0 0 0 2.336-2.342v0A2.339 2.339 0 0 0 25.026 23H23.44a2.339 2.339 0 0 1-2.337-2.341v0a2.339 2.339 0 0 1 2.337-2.342h.793a3.975 3.975 0 0 1 2.907.79m9.36 6.796a3.533 3.533 0 0 1-3.069 1.78h0a3.534 3.534 0 0 1-3.534-3.534V21.85a3.534 3.534 0 0 1 3.534-3.534h0a3.533 3.533 0 0 1 3.066 1.774"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}