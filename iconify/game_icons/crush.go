package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M46 16c-15 0-30 15-30 30v240h60l60 60l60-60l60 60l60-60l60 60l60-60h60V46c0-15-15-30-30-30H46zm211.75 360.03A60 60 0 0 0 196 436a60 60 0 0 0 120 0a60 60 0 0 0-58.25-59.97z"/>`),
		g.Group(children),
	)
}