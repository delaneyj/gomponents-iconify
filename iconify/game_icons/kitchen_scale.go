package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenScale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m50.73 121l49.57 62h311.4l49.6-62H50.73zM245 201v46h22v-46h-22zm-126 64L75.53 439H436.5L393 265H119zm137 14c40.2 0 73 32.8 73 73s-32.8 73-73 73s-73-32.8-73-73s32.8-73 73-73zm0 18c-30.5 0-55 24.5-55 55s24.5 55 55 55s55-24.5 55-55s-24.5-55-55-55zm16 12.5l-5.9 65.7l-30.2-10.5l36.1-55.2zM41 457v30h430v-30H41z"/>`),
		g.Group(children),
	)
}