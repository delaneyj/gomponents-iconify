package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videogame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 360h480V120H16Zm32-208h416v176H48Z"/><path fill="currentColor" d="M152 184h-32v40H80v32h40v40h32v-40h40v-32h-40v-40zm184 72h32v32h-32zm64-64h32v32h-32z"/>`),
		g.Group(children),
	)
}