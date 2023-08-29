package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HackerNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M0 32v448h448V32H0zm21.2 197.2H21c.1-.1.2-.3.3-.4c0 .1 0 .3-.1.4zm218 53.9V384h-31.4V281.3L128 128h37.3c52.5 98.3 49.2 101.2 59.3 125.6c12.3-27 5.8-24.4 60.6-125.6H320l-80.8 155.1z"/>`),
		g.Group(children),
	)
}