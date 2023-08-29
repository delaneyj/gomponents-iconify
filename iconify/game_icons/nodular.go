package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nodular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16a90 90 0 0 0-15 178.593v48.75a60 60 0 0 0 0 115.782v49.688A45 45 0 0 0 256 496a45 45 0 0 0 15-87.188v-50.157a60 60 0 0 0 0-115.78V194.59a90 90 0 0 0-15-178.594zm-2.813 30A60 60 0 0 1 256 46a60 60 0 0 1 0 120a60.017 60.017 0 0 1-2.813-120z"/>`),
		g.Group(children),
	)
}