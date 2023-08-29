package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaugeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M205.65 82.08A109.26 109.26 0 0 0 128 50h-.39C67.17 50.21 18 100.06 18 161.13V184a14 14 0 0 0 14 14h192a14 14 0 0 0 14-14v-24a109.3 109.3 0 0 0-32.35-77.92ZM226 184a2 2 0 0 1-2 2H115.78l57.07-78.47a6 6 0 0 0-9.7-7.06L100.94 186H32a2 2 0 0 1-2-2v-22.87a102.36 102.36 0 0 1 .62-11.13H56a6 6 0 0 0 0-12H32.71C42.6 96.4 78.53 64.86 122 62.19V88a6 6 0 0 0 12 0V62.19A98.05 98.05 0 0 1 223.53 138H200a6 6 0 0 0 0 12h25.5c.33 3.3.5 6.64.5 10Z"/>`),
		g.Group(children),
	)
}