package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M105 41v398h302V41H105zm55 174c18.1 0 33 14.9 33 33s-14.9 33-33 33s-33-14.9-33-33s14.9-33 33-33zm0 18c-8.4 0-15 6.6-15 15s6.6 15 15 15s15-6.6 15-15s-6.6-15-15-15zM73 457v30h366v-30H73z"/>`),
		g.Group(children),
	)
}