package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 152a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-4.17 35.79A44.06 44.06 0 0 1 160 228h-24a44.05 44.05 0 0 1-44-44v-44.15A60 60 0 0 1 36 80V40a12 12 0 0 1 12-12h16a4 4 0 0 1 0 8H48a4 4 0 0 0-4 4v40a52 52 0 0 0 52.69 52C125 131.63 148 107.93 148 79.17V40a4 4 0 0 0-4-4h-16a4 4 0 0 1 0-8h16a12 12 0 0 1 12 12v39.17c0 32-24.84 58.59-56 60.69V184a36 36 0 0 0 36 36h24a36.05 36.05 0 0 0 35.8-32.25a36 36 0 1 1 8 0ZM228 152a28 28 0 1 0-28 28a28 28 0 0 0 28-28Z"/>`),
		g.Group(children),
	)
}