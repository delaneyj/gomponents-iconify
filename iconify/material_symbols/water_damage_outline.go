package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterDamageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-8H2l10-9l10 9h-3v8H5Zm2-2h10v-7.8l-5-4.5l-5 4.5V18Zm5-2q.825 0 1.413-.588T14 14q0-.675-.375-1.438T12 10q-1.25 1.8-1.625 2.563T10 14q0 .825.588 1.413T12 16Zm0-4.15Z"/>`),
		g.Group(children),
	)
}