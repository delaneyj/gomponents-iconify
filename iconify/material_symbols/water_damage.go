package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterDamage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-8H2l10-9l10 9h-3v8H5Zm7-4q.825 0 1.413-.588T14 14q0-.675-.375-1.438T12 10q-1.25 1.8-1.625 2.563T10 14q0 .825.588 1.413T12 16Z"/>`),
		g.Group(children),
	)
}