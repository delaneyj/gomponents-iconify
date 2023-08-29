package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M41 0h525c23 0 41 18 41 41v592c0 23-18 41-41 41H431V541H176v133H41c-23 0-41-18-41-41V41C0 18 18 0 41 0zm49 134h427V84H90v50zm0 115h427v-50H90v50zm0 115h427v-50H90v50zm0 115h427v-50H90v50zm200 217h-87V568h87v128zm114 0h-87V568h87v128z"/>`),
		g.Group(children),
	)
}