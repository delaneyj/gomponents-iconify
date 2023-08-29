package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenPointedStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 22.017l-69.427 102.007l-123.038-9.32L100 232.584l-84 90.384l114.898 44.987l18.292 122.028L256 428.2l106.81 61.783l18.292-122.028L496 322.968l-84-90.385l36.465-117.88l-123.038 9.32z"/>`),
		g.Group(children),
	)
}