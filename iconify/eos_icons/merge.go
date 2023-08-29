package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Merge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 11a2.991 2.991 0 0 0-2.815 1.997c-2.995-.014-5.26-.129-6.726-.886a6.205 6.205 0 0 1-3.358-4.326A3.008 3.008 0 1 0 4 7.816v9.368a3 3 0 1 0 2 0v-5.257a8.579 8.579 0 0 0 2.541 1.962c1.847.952 4.36 1.092 7.642 1.108A2.995 2.995 0 1 0 19 11Z"/>`),
		g.Group(children),
	)
}