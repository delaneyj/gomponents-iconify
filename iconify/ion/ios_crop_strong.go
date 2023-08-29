package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosCropStrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M128 64h32v48h-32z" fill="currentColor"/><path d="M160 352V176h-32v208h208v-32z" fill="currentColor"/><path d="M400 352h48v32h-48z" fill="currentColor"/><path d="M64 128v32h288v288h32V128z" fill="currentColor"/>`),
		g.Group(children),
	)
}