package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlainCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 23.05C127.5 23.05 23.05 127.5 23.05 256S127.5 488.9 256 488.9S488.9 384.5 488.9 256S384.5 23.05 256 23.05z"/>`),
		g.Group(children),
	)
}