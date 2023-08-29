package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iraq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m17.03 246.5l25.73 60.7l265.44 143l82.2 6.9l38.1-47.1l66.5 4.1l-33.5-52.1l-5.2-38l-116.5-107.9l39-53.8l-72.3-100.65l-95.4-6.77l-69.5 49.82l-16.4 97.2z"/>`),
		g.Group(children),
	)
}