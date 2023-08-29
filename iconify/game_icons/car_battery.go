package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M73 81v46h46V81H73zm320 0v46h46V81h-46zM25 145v16h462v-16H25zm0 34v252h462V179H25zm382 21h18v23h23v18h-23v23h-18v-23h-23v-18h23v-23zM64 223h64v18H64v-18z"/>`),
		g.Group(children),
	)
}