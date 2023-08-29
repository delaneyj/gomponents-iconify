package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Knockonports(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.06 5.57h25.88a.83.83 0 0 1 .66 1.33L29.29 18a.84.84 0 0 1-.67.33h-9.24a.81.81 0 0 1-.66-.33L10.4 6.9a.83.83 0 0 1 .66-1.33Z"/><circle cx="24" cy="38.74" r="5.83" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.8 38.22A14.28 14.28 0 0 0 34.05 15M14 15a14.28 14.28 0 0 0 4.2 23.22"/>`),
		g.Group(children),
	)
}