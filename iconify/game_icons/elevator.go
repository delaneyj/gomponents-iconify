package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M153 35v58h206V35H153zm60.3 13h32l-16 32l-16-32zm74.7 0l16 32h-32l16-32zm-183 89v350h142V137H105zm160 0v350h142V137H265zm173 141v84h52v-84h-52zm26 26a16 16 0 0 1 16 16a16 16 0 0 1-16 16a16 16 0 0 1-16-16a16 16 0 0 1 16-16z"/>`),
		g.Group(children),
	)
}