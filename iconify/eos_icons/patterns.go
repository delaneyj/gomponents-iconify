package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patterns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="4.02" cy="4.015" r="3" fill="currentColor"/><path fill="currentColor" d="M15.03 12.025a2.991 2.991 0 0 0-2-2.816v-2.38a2.99 2.99 0 0 0 1.807-1.814h2.367a2.99 2.99 0 0 0 1.816 1.816v2.38a3.001 3.001 0 1 0 2-.005V6.83a2.993 2.993 0 1 0-3.816-3.816h-2.367a2.993 2.993 0 1 0-3.807 3.82v2.374a2.983 2.983 0 0 0 0 5.632v2.375a2.988 2.988 0 0 0-1.827 1.823H6.838a2.989 2.989 0 0 0-1.818-1.82v-2.375a2.999 2.999 0 1 0-2-.006v2.38a2.993 2.993 0 1 0 3.815 3.821h2.37a2.993 2.993 0 1 0 3.825-3.817v-2.38a2.991 2.991 0 0 0 2-2.817Z"/><circle cx="20.02" cy="20.035" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}