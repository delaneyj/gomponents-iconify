package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdStats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M176 64h64v384h-64z" fill="currentColor"/><path d="M80 336h64v112H80z" fill="currentColor"/><path d="M272 272h64v176h-64z" fill="currentColor"/><path d="M368 176h64v272h-64z" fill="currentColor"/>`),
		g.Group(children),
	)
}