package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M466 0H251c-26 0-45 20-45 46v250H36c-16 0-27 7-34 22c-1 5-2 9-2 12c0 10 3 19 10 26l324 323c12 15 36 14 50 0l323-323c22-20 7-60-26-60H510V46c0-26-18-46-44-46z"/>`),
		g.Group(children),
	)
}