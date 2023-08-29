package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M134.5 30.5v451h243v-451zm100.68 20h41.6a8 8 0 0 1 0 16h-41.6a8 8 0 1 1 0-16zm20.32 420.51a19.26 19.26 0 1 1 19.26-19.26a19.26 19.26 0 0 1-19.26 19.26zm105-44.51h-211v-343h211z"/>`),
		g.Group(children),
	)
}