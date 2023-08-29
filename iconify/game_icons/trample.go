package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trample(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M136 16h15l15 15l15-15h150l15 15l15-15h15v120h-15l-15-15l-15 15h-45v345c0 15-15 15-15 15h-30s-15 0-15-15V136h-45l-15-15l-15 15h-15z"/>`),
		g.Group(children),
	)
}