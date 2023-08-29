package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Venus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1152 576q0 221-147.5 384.5T640 1148v260h224q14 0 23 9t9 23v64q0 14-9 23t-23 9H640v224q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-224H288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h224v-260q-150-16-271.5-103t-186-224T2 529q11-134 80.5-249t182-188T510 4q170-19 319 54t236 212t87 306zm-1024 0q0 185 131.5 316.5T576 1024t316.5-131.5T1024 576T892.5 259.5T576 128T259.5 259.5T128 576z"/>`),
		g.Group(children),
	)
}