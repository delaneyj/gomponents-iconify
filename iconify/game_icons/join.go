package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Join(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 23v18h215v110h-37.6l6.5 13l40.1 80.1l46.6-93.1H265V41h215V23zm224 244.9L209.4 361H247v110H32v18h448v-18H265V361h37.6z"/>`),
		g.Group(children),
	)
}