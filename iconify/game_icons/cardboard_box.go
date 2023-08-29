package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardboardBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M185.424 52.607L30.404 130.04l55.375 55.317l155.017-77.43l-55.373-55.32zm141.152 0l-55.373 55.32l155.018 77.43l55.376-55.316l-155.02-77.433zM256 120.45l-9 4.497v142.715l9 4.496l9-4.496V124.947l-9-4.496zM86.482 207.605l-57.59 71.917l139.545 77.45l72.358-72.286l-154.313-77.08zm339.036 0l-154.313 77.08l72.358 72.287l139.544-77.45l-57.59-71.916zM247 303.93l-75.436 75.36l-78.562-43.6v44.058L247 456.67V303.93zm18 0v152.74l153.998-76.922v-44.06l-78.562 43.603L265 303.93z"/>`),
		g.Group(children),
	)
}