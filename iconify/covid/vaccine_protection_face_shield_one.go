package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineProtectionFaceShieldOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 .81c-3.559.543-6 1.645-6 2.92c0 1.811 4.925 3.28 11 3.28s11-1.469 11-3.28c0-1.275-2.441-2.377-6-2.92"/><path d="M1 3.73v3.3c0 1.812 4.925 3.281 11 3.281s11-1.465 11-3.277V3.73"/><path d="m4 9.315l-.9 9.794a1.98 1.98 0 0 0 .848 1.823A14.277 14.277 0 0 0 12 23.19c2.85.068 5.655-.718 8.054-2.258a1.98 1.98 0 0 0 .848-1.823L20 9.315"/><path d="M13 20.19a7.638 7.638 0 0 0 5-2"/></g>`),
		g.Group(children),
	)
}