package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M204 181.372L38.628 16H16v22.628L181.372 204H44v32h192V44h-32v137.372zM326.628 304H464v-32H272v192h32V326.628L473.372 496H496v-22.628L326.628 304z"/>`),
		g.Group(children),
	)
}