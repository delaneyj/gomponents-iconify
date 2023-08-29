package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480 40L32 296l112.148 37.383L448 72L209.404 355.135L320 392L480 40zM208 376l-16 96l49.932-83.863L208 376z"/>`),
		g.Group(children),
	)
}