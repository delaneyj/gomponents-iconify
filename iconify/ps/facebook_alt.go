package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 78q0-30-23-53T386 2H78Q48 2 25 25T2 78v308q0 30 23 53t53 23h154V288h-56v-76h56v-30q0-39 25.5-68.5T318 84h62v76h-62q-5 0-9.5 6t-4.5 15v31h76v76h-76v174h82q30 0 53-23t23-53V78z"/>`),
		g.Group(children),
	)
}