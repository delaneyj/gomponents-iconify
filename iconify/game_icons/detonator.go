package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Detonator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M25 41v46h94V41H25zm368 0v46h94V41h-94zM137 55v18h110v110h18V73h110V55H137zM73 201v30h366v-30H73zm32 48v190h302V249H105zm151 17l96 150H160l96-150zm-9 38v64h18v-64h-18zm9 75a12 12 0 0 0-12 12a12 12 0 0 0 12 12a12 12 0 0 0 12-12a12 12 0 0 0-12-12zM73 457v30h366v-30H73z"/>`),
		g.Group(children),
	)
}