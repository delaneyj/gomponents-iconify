package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ladder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M121 17v30h270V17H121zm16 48v46h30V65h-30zm208 0v46h30V65h-30zm-224 64v30h270v-30H121zm16 48v46h30v-46h-30zm208 0v46h30v-46h-30zm-224 64v30h270v-30H121zm16 48v46h30v-46h-30zm208 0v46h30v-46h-30zm-224 64v30h270v-30H121zm16 48v46h30v-46h-30zm208 0v46h30v-46h-30zm-224 64v30h270v-30H121z"/>`),
		g.Group(children),
	)
}