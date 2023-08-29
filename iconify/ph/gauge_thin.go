package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaugeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204.23 83.5A107.37 107.37 0 0 0 127.62 52C68.28 52.21 20 101.16 20 161.13V184a12 12 0 0 0 12 12h192a12 12 0 0 0 12-12v-24a107.25 107.25 0 0 0-31.77-76.5ZM228 184a4 4 0 0 1-4 4H111.85l59.38-81.65a4 4 0 1 0-6.46-4.7L102 188H32a4 4 0 0 1-4-4v-22.87a103.42 103.42 0 0 1 .84-13.13H56a4 4 0 0 0 0-8H30.21C39.59 95.66 77.84 61.93 124 60.09V88a4 4 0 0 0 8 0V60.08A100.08 100.08 0 0 1 226 140h-26a4 4 0 0 0 0 8h27.29a101.6 101.6 0 0 1 .71 12Z"/>`),
		g.Group(children),
	)
}