package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M67.79 62.74a4 4 0 0 0-7.58 0l-40 120A4 4 0 0 0 24 188h80a4 4 0 0 0 3.79-5.26ZM29.55 180L64 76.65L98.45 180ZM204 76a48 48 0 1 0-48 48a48.05 48.05 0 0 0 48-48Zm-88 0a40 40 0 1 1 40 40a40 40 0 0 1-40-40Zm108 72h-88a4 4 0 0 0-4 4v56a4 4 0 0 0 4 4h88a4 4 0 0 0 4-4v-56a4 4 0 0 0-4-4Zm-4 56h-80v-48h80Z"/>`),
		g.Group(children),
	)
}