package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 30v256h-64l96 128l96-128h-64V30h-64zM32 434v48h448v-48H32z"/>`),
		g.Group(children),
	)
}