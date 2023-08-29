package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HadesSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 16L112 272l72.8 53l55-37.1L192 256zm32 0l48 240l-144 96l-64 48l144 96l144-96l-64-48l-8.2-5.4l-56.5 41.2L288 400l-32 32l-32-32l176-128z"/>`),
		g.Group(children),
	)
}