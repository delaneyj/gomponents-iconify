package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosPauseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M191 112v288h-47V112h47m16-16h-79v320h79V96z" fill="currentColor"/><path d="M368 112v288h-47V112h47m16-16h-79v320h79V96z" fill="currentColor"/>`),
		g.Group(children),
	)
}