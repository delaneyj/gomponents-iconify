package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="m0 32l34.9 395.8L191.5 480l157.6-52.2L384 32H0zm308.2 127.9H124.4l4.1 49.4h175.6l-13.6 148.4l-97.9 27v.3h-1.1l-98.7-27.3l-6-75.8h47.7L138 320l53.5 14.5l53.7-14.5l6-62.2H84.3L71.5 112.2h241.1l-4.4 47.7z"/>`),
		g.Group(children),
	)
}