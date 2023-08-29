package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pounce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m196 16l60 60l60-60H196zm-30 90l90 90l90-90H166zm-30 120l120 120l120-120H136zm121.75 150.03A60 60 0 0 0 196 436a60 60 0 0 0 120 0a60 60 0 0 0-58.25-59.97z"/>`),
		g.Group(children),
	)
}