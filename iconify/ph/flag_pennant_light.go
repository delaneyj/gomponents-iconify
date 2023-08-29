package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPennantLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m242 98.33l-184-64A6 6 0 0 0 50 40v176a6 6 0 0 0 12 0v-43.73l180-62.6a6 6 0 0 0 0-11.34ZM62 159.56V48.44L221.74 104Z"/>`),
		g.Group(children),
	)
}