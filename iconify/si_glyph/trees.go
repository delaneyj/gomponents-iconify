package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trees(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.986 14.203v-2.278c1.52-.339 2.967-1.952 2.967-3.896c0-2.19-1.648-7.982-3.451-7.982c-1.803 0-3.455 5.792-3.455 7.982c0 1.901 1.513 3.489 2.984 3.874v2.21a62.349 62.349 0 0 0-3.011-.072a24.06 24.06 0 0 0-3.067.197v-2.672c1.058-.311 1.97-1.418 1.97-2.75c0-1.57-1.095-5.781-2.442-5.781c-1.35 0-2.442 4.211-2.442 5.781c0 1.354.905 2.48 1.993 2.769v2.793c-2.395.413-4.01 1.113-4.01 1.567h16c-.001-.456-1.627-1.531-4.036-1.742z"/>`),
		g.Group(children),
	)
}