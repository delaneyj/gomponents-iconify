package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlainArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M130.81 21.785v245.95H43.84L256 489.382l212.158-221.644H381.19V21.786H130.81z"/>`),
		g.Group(children),
	)
}