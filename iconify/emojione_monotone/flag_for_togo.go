package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM10.322 49.699a27.911 27.911 0 0 1-5.691-11.8h54.738a27.923 27.923 0 0 1-5.691 11.8H10.322zM24.919 32l-6.686-4.616L11.548 32l2.523-7.505l-6.654-4.651l8.243-.022l2.573-7.488l2.571 7.488l8.245.022l-6.66 4.651L24.919 32zM32 26.1V14.3h21.677a27.913 27.913 0 0 1 5.692 11.8H32z"/>`),
		g.Group(children),
	)
}